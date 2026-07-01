# Signature Verification

Verifikasi signature adalah salah satu teknik yang digunakan untuk mengindentifikasi apakah sebuah request berasal dari sumber yang sah atau tidak. 

## Contoh use case
**Payment callback pada ecommerce**
1. User melakukan create order.
2. System membuat order dan payment dengan status pending yang akan expire dalam 5 menit.
3. User diarahkan ke halaman milik payment gateway.
4. User melakukan pembayaran.
5. Payment gateway mengirim payment callback request ke system.

Ada beberapa teknik untuk memastikan bahwa callback request ini valid, salah satunya adalah memastikan bahwa payload "di-tandai" dengan signature yang telah disepakati oleh system ecommerce dan system payment gateway. Ini disebut dengan teknik signature verification.

Ada dua cara:
1. HMAC signature.
2. public key signature.

### HMAC Signature
HMAC -> Hash-based Message Authentication Code

1. Hanya menggunakan 1 secret key. Client dan server menggunakan secret key yang sama.
2. Client:
   - membuat signature menggunakan secret key dan menandai request body dengan signature tersebut. 
   - Client mengirim signature lewat request header, contoh: `X-Signature`
3. Saat menerima request, server melakukan:
   - membuat expected signature menggunakan request body dan secret key yang sama.
   - membandingkan expected signature dengan signture yang dikirim oleh client.
   - jika signature sama -> proses.
   - jika beda -> reject.

#### contoh:

Client:
```
func sendCallback(paymentID string, status string) error {
	payload := fmt.Sprintf(`{
		"payment_id": "%s",
		"status" :"%s"
	}`, paymentID, status)

	// create signature
	sign := GenerateHMAC(payload, "ini-rahasia")

	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080/payments/callback",
		strings.NewReader(payload),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Signature", sign)

    // rest of the code
    ...

}

func GenerateHMAC(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
```

Server:
```
func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return helper.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	signature := r.Header.Get("X-Signature")

	ctx := context.Background()
	secret = "ini-rahasia"

	ok := helper.VerifyHMAC(string(body), secret, signature)
	if !ok {
		err := errors.New("Wrong signature")
		return helper.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))
    
    // rest of the code
    ...
}

func VerifyHMAC(message, secret, givenSignature string) bool {
	expectedHMAC := hmac.New(sha256.New, []byte(secret))
	expectedHMAC.Write([]byte(message))
	expectedSignature := expectedHMAC.Sum(nil)

	decodedGiven, err := hex.DecodeString(givenSignature)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedSignature, decodedGiven)
}
```