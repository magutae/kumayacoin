package wallet

// 1) hash the msg
// "message" -> hash(x) -> "hashed_message"

// 2) generate key pair
// KeyPair (privateKey, publicKey) -> (save priv to a file: wallet)

// 3) sign the hash
// ("hashed_message" + privateKey) -> "signature"

// 4) verify
// ("hashed_message", + "segnature" + publicKey) -> true / false

const (
	signature     string = "18bcb6862fdf459c33021091f521fd37a74d62006488df2fde65d79bbcdf98ba85601f364fbc3942b2cc1961ef4b866e4bfccbfab51b422473991f482b6b3adb"
	privateKey    string = "30770201010420cca45b6a3f0209b12e5c3f67c43fe2c9b0555e4b7a1f4b32ff1a8d7680359564a00a06082a8648ce3d030107a14403420004db99c4dc07d960b2793745290be09be23be58fab275a9b5018b33fbc15d73a18d8635d9f189db5a9d0915b0a6a1a9ad663eea5819eb44f1ce65703f2f8e48ae3"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {

}
