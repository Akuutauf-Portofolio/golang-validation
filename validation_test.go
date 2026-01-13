package belajar_go_lang_validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

// implementasi dari validation struct
// membuat pengujian untuk validation struct
func TestValidation(t *testing.T) {
	// membuat object validate baru, singleton: cukup sekali buat saja untuk seluruh aplikasi
	// karena validator menggunakan struct, maka tambahkan pointer
	var validate *validator.Validate = validator.New()

	// menerapkan object validate
	// akan mengecek apakah variabel validate kosong?
	if validate == nil {
		// kalau kosong akan menampilkan pesan berikut
		t.Error("Validate is nil")
	}
}

// implementasi validation variable
// membuat pengujian untuk validation variable menggunakan tag
func TestValidationVariable(t *testing.T) {
	// membuat object validate baru
	validate := validator.New()

	// membuat variabel baru
	text := "taufik"

	// melakukan pengecekan validasi variabel dengan menggunakan tag 'required'
	// tag required digunakan untuk melakukan pengecekan jika variabel tidak memiliki value, maka akan error
	// untuk melakukan validasi di variabel menggunakan function validate.Var / validate.VarCtx
	err := validate.Var(text, "required")

	// mengecek hasil validasi
	// jika error tidak ada maka akan berhasil
	if err != nil {
		// jika kosong, maka akan menampilkan error
		fmt.Println(err.Error()) // akan error jika variabel kosong
	} else {
		fmt.Println("Tidak ada error")
	}	
}

// implementasi validation two variable
func TestValidateTwoVariable(t *testing.T) {
	// membuat object validate
	validate := validator.New()

	// menyiapkan 2 buah variabel
	password := "rahasia123"
	confirmPassword := "secret123"

	// melakukan validation dengan mengunakan tag 'eqfield'
	// Validate.VarWithValue / Validate.VarWithValueCtx, digunakan untuk membandingkan isi dari dua buah variabel
	// tag 'eqfield' ini digunakan untuk membandingkan nilai dari kedua variabel sama
	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	// melakukan pengecekan
	if err != nil {	
		fmt.Println("Konfirmasi password gagal")
		fmt.Print(err.Error())
	} else {
		fmt.Println("Konfirmasi password sukses")
	}
}

// implementasi multiple tag (lebih dari satu tag)
func TestMultipleTag(t *testing.T) {
	// membuat object validate baru
	validate := validator.New()

	// membuat variabel baru
	text := "12345"

	// bisa melakukan validasi dengan tag lebih dari satu, dengan menggunakan koma (,) sebagai pemisah
	// tag required : wajib di isi nilai variabel nya
	// tag number : nilai variabel wajib berisi hanya angka
	err := validate.Var(text, "required,number")

	// mengecek hasil validasi
	// jika error tidak ada maka akan berhasil
	if err != nil {
		// jika kosong, maka akan menampilkan error
		// akan error jika variabel bernilai nill dan bukan selain angka atau kosong
		fmt.Println(err.Error()) 
	} else {
		fmt.Println("Tidak ada error")
	}	
}

// implementasi tag parameter
func TestTagParameter(t *testing.T) {
	// membuat object validate baru
	validate := validator.New()

	// membuat variabel baru
	text := "12345"

	// validator bisa melakukan tag parameter dengan memberikan sama dengan '=' setelah nama tag
	err := validate.Var(text, "required,number,min=5,max=10")

	// mengecek hasil validasi
	// jika error tidak ada maka akan berhasil
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	} else {
		fmt.Println("Tidak ada error")
	}	
}

// implementasi validasi struct
func TestStructValidate(t *testing.T) {
	// membuat struct baru
	type LoginRequest struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	loginUser := LoginRequest{
		Username: "taufik@gmail.com",
		Password: "taufik",
	}

	// nah sekarang untuk melakukan validasi tidak lagi mengunakan Validate.Var, melainkan-
	// menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	err := validate.Struct(loginUser)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}

// implementasi validation errors
func TestValidationErrors(t *testing.T) {
	// membuat struct baru
	type LoginRequest struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	loginUser := LoginRequest{
		Username: "taufik",
		Password: "123",
	}

	// nah sekarang untuk melakukan validasi tidak lagi mengunakan Validate.Var, melainkan-
	// menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	err := validate.Struct(loginUser)

	// mengecek hasil validasi
	if err != nil {
		// jika validasi error
		// melakukan konversi error menjadi ValidationErrors (function di package validator)
		validationErrors := err.(validator.ValidationErrors)

		// menampilkan beberapa informasi erorr yang sudah dikonversi
		for _, fieldError := range validationErrors {
			// bisa menampilkan seluruh informasi error dengan detail sesuai dengan-
			// informasi yang ada pada []FieldError
			// function fieldError.Field() : digunakan untuk mengambil informasi field yang error
			// function fieldError.Tag() : digunakan untuk mengambil informasi tag yang menyebabkan error
			// function fieldError.Error() : digunakan untuk mengambil pesan errornya seluruhnya
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

// implementasi cross field pada struct
func TestStructCrossField(t *testing.T) {
	// membuat struct baru
	type RegisterUser struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,eqfield=Password"` // menambahkan field yang ingin dikonfirmasi (dibandingkan)
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	request := RegisterUser{
		Username: "taufik@gmail.com",
		Password: "taufik123",
		ConfirmPassword: "taufik123",
	}

	// nah sekarang untuk melakukan validasi tidak lagi mengunakan Validate.Var, melainkan-
	// menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	err := validate.Struct(request)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}