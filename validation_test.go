package belajar_go_lang_validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

// implementasi validasi nested struct
func TestNestedStruct(t *testing.T) {
	// membuat struct baru (pertama)
	type Address struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	// membuat struct berikutnya (kedua)
	type User struct {
		// memiliki attribute seperti biasanya (normal)
		Id string `validate:"required"`
		Name string `validate:"required"`

		// namun di sini kita tambahkan attribute yang menggunakan struct pertama
		Address Address `validate:"required"`
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	request := User{
		Id: "1",
		Name: "Taufik",
		Address: Address{
			City: "Banyuwangi",
			Country: "Indonesia",
		},
	}

	// validasi menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	
	// pada nested struct secara otomatis nanti attribute yang menggunakan struct pertama, akan dilakukan validasi-
	// didalamnya, asalkan sudah menambahkan validasi di attribute pada struct pertama dengan metode relfection
	err := validate.Struct(request)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}

// implementasi validasi collection
func TestCollection(t *testing.T) {
	// membuat struct baru (pertama)
	type Address struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	// membuat struct berikutnya (kedua)
	type User struct {
		// memiliki attribute seperti biasanya (normal)
		Id string `validate:"required"`
		Name string `validate:"required"`

		// namun di sini kita tambahkan attribute yang menggunakan struct pertama
		// dengan menjadikan sebagai collection, berupa slice
		Address []Address `validate:"required,dive"`

		// secara default validator package tidak bisa memvalidasi isi dari attribute struct yang-
		// berjenis collection, maka dari itu jika butuh untuk validasi isi dari slice. tambahkan-
		// tag yang bernama 'dive'
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	request := User{
		Id: "1",
		Name: "Taufik",
		Address: []Address{
			Address{
				City: "",
				Country: "",
			},
			Address{
				City: "Surabaya",
				Country: "Indonesia",
			},
		},
	}

	// validasi menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	
	// pada nested struct secara otomatis nanti attribute yang menggunakan struct pertama, akan dilakukan validasi-
	// didalamnya, asalkan sudah menambahkan validasi di attribute pada struct pertama dengan metode relfection
	err := validate.Struct(request)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}

// implementasi validasi basic collection
func TestBasicCollection(t *testing.T) {
	// membuat struct baru (pertama)
	type Address struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	// membuat struct berikutnya (kedua)
	type User struct {
		// memiliki attribute seperti biasanya (normal)
		Id string `validate:"required"`
		Name string `validate:"required"`
		Address []Address `validate:"required,dive"`

		// menambahkan attribute baru yang masih collection, namun tipe data nya selain struct
		Hobbies []string `validate:"required,dive,required,min=3"`
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	request := User{
		Id: "1",
		Name: "Taufik",
		Address: []Address{
			Address{
				City: "Banyuwangi",
				Country: "Indonesia",
			},
			Address{
				City: "Surabaya",
				Country: "Indonesia",
			},
		},
		Hobbies: []string {
			"Reading", "Gaming", "X", "",

			// value hobbies yang x dan (kosong), akan mengembalikan error validasi
			// karena nilai minimal untuk setiap data hobbies haruslah 3 karakter
		},
	}

	// validasi menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	err := validate.Struct(request)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}

// implementasi validasi map
func TestMap(t *testing.T) {
	// membuat struct baru (pertama)
	type Address struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	// membuat strut selanjutnya (ketiga)
	type School struct {
		Name string `validate:"required"`
	}

	// membuat struct berikutnya (kedua)
	type User struct {
		// memiliki attribute seperti biasanya (normal)
		Id string `validate:"required"`
		Name string `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string `validate:"required,dive,required,min=3"`

		// menambahkan attribute baru berupa map,-
		// dengan keys nya adalah string dan value nya adalah object struct school
		Schools map[string]School `validate:"required,min=1,dive,keys,required,min=2,endkeys"`

		// tag required diawal : agar attribute schools di isi (tidak boleh kosong)
		// tag dive pertama : digunakan agar pengecekan nya bisa lebih dalam untuk validasinya
		// tag keys dan enkeys : di antara 2 tag ini di dalamnya terdapat tag yang diperuntukkan khusus keys
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	request := User{
		Id: "1",
		Name: "Taufik",
		Address: []Address{
			Address{
				City: "Banyuwangi",
				Country: "Indonesia",
			},
			Address{
				City: "Surabaya",
				Country: "Indonesia",
			},
		},
		Hobbies: []string {
			"Reading", "Gaming", "Ngoding", "Running",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD 1 Pancasila",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
	}

	// validasi menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	err := validate.Struct(request)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}

// implementasi validasi basic map
func TestBasicMap(t *testing.T) {
	// membuat struct baru (pertama)
	type Address struct {
		// memiliki attribute yang mengimplementasikan validasi, dengan konsep relfection
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	// membuat strut selanjutnya (ketiga)
	type School struct {
		Name string `validate:"required"`
	}

	// membuat struct berikutnya (kedua)
	type User struct {
		// memiliki attribute seperti biasanya (normal)
		Id string `validate:"required"`
		Name string `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string `validate:"required,dive,required,min=3"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys"`

		// menambahkan attribute baru, untuk map basic dengan tipe data yang biasa
		Wallets map[string]int `validate:"dive,keys,required,endkeys,required,gt=1000"`

		// tag dive : digunakan untuk validasi agar masuk ke dalam key dan valuenya
		// tag keys dan endkeys : digunakan untuk validasi keys nya saja
		// dan tag setelah enkeys : adalah validasi untuk bagian value
	}

	// membuat object validate baru
	validate := validator.New()

	// membuat object baru dari struct sebelumnya
	request := User{
		Id: "1",
		Name: "Taufik",
		Address: []Address{
			Address{
				City: "Banyuwangi",
				Country: "Indonesia",
			},
			Address{
				City: "Surabaya",
				Country: "Indonesia",
			},
		},
		Hobbies: []string {
			"Reading", "Gaming", "Ngoding", "Running",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD 1 Pancasila",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
		Wallets: map[string]int{
			"BNI": 100000,
			"BCA": 500000,
			"BRI": 0,
			"MANDIRI": 1000,
		},
	}

	// validasi menggunakan function Validate.Struct(), karena pada struct sudah di berikan validasinya
	err := validate.Struct(request)

	// mengecek hasil validasi
	if err != nil {
		// akan menampilkan erorr jika salah satu validasi tidak terpenuhi
		fmt.Println(err.Error()) 
	}
}

// implementasi alias
func TestAlias(t *testing.T) {
	// membuat object validate baru
	validate := validator.New()

	// membuat alias untuk beberapa tag menjadi tag baru dengan nama custom
	validate.RegisterAlias("varchar", "required,max=255") // bisa lebih dari satu tag yang di register di alias baru

	// membuat struct yang mengimplementasikan alias
	// maka akan secara otomatis, akan mengimplementasikan seluruh validasi tag yang di dafatarkan-
	// pada tag alias 'varchar'
	type Seller struct {
		Id string `validate:"varchar,min=5"`
		Name string `validate:"varchar"`
		Owner string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	// membuat object dari struct
	seller := Seller {
		Id: "123",
		Name: "",
		Owner: "",
		Slogan: "",
	}

	// melakukan validasi pada object struct 
	err := validate.Struct(seller)

	// mengecek error validasi
	if err != nil {
		// jika terdapat error, akan ditampilkan
		fmt.Println(err.Error())
	}
}

// implementasi custom validation

// membuat function custom untuk di register nantinya
// parameternya wajib menggunakan field level (bawaan validator)
func MustValidUsername(field validator.FieldLevel) bool {
	// mengambil data dari relfection, dan mengkonversi dalam bentuk interface lalu ke string
	value, ok := field.Field().Interface().(string)

	// mengecek jika sukses data yang dimbil
	if ok {
		// mengecek jika data value tidak huruf besar semua, maka tidak valid
		if value != strings.ToUpper(value) {
			return false
		} 

		// mengecek jika data value panjang nya kurang dari 5, maka tidak valid
		if len(value) < 5 {
			return false
		}
	}

	// kalau sukses semua pengecekannya, bisa return true (artinya valid usernamenya)
	return true
}

func TestCustomValidationFunction(t *testing.T) {
	// membuat object validate
	validate := validator.New()

	// melakukan register validation custom dari function yang sudah dibuat sebelumnya
	validate.RegisterValidation("username", MustValidUsername)

	// membuat struct baru
	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	// membuat object baru dari struct
	request := LoginRequest{
		Username: "AKUUTAUF",
		Password: "rahasia",
	}

	// melakukan validasi dari struct
	err := validate.Struct(request)

	// mengecek error validasi
	if err != nil {
		fmt.Println(err.Error())
	}
}

// implementasi custom validation parameter

// membuat aturan regex untuk pin
// ^[0-9] : jumlah karakter bebas apapun itu, yang penting berkisar antara 0 hingga 9
var regexNumber = regexp.MustCompile("^[0-9]+$")

// membuat function custom baru
func MustValidPin(field validator.FieldLevel) bool {
	// mengambil data field level, hanya bagian field.param
	// return valude berupa string, meski inputnya adalah angka. namun tetap string untuk outputnya
	// maka dari itu perlu dikonversi dari string ke bentuk yang di inginkan, misal ke integer (menggunakan Atoi)
	length, err := strconv.Atoi(field.Param())

	// mengecek jika input data tidak valid
	if err != nil {
		panic(err)
	}

	// mengambil value dengan mengkonversi ke bentuk interface lalu string
	value := field.Field().Interface().(string)

	// melakukan validasi pin harus regex, dengan mengecek inputnya match string atau tidak
	if !regexNumber.MatchString(value) {
		// jika tidak match input dengan regexp yang sudah dibuat, maka akan return false
		return false
	}

	// kalau cocok dengan regexp, selanjutnya mengecek panjang inputnya
	// jika benar semua pengecekannya, return panjangnya,-
	// yang mana panjang value harus sama dengan panjang parameter
	return len(value) == length // kalau salah akan menghasilkan false, dan sebaliknya
}

func TestCustomValidationFunctionWithParameter(t *testing.T) {
	// membuat object validate
	validate := validator.New()

	// melakukan register validation custom dari function yang sudah dibuat sebelumnya
	validate.RegisterValidation("pin", MustValidPin)

	// membuat struct baru
	type LoginRequest struct {
		Username string `validate:"required,number"`
		Password string `validate:"required,pin=6"`
	}

	// membuat object baru dari struct
	request := LoginRequest{
		Username: "12345",
		Password: "12345", // akan error jika panjang password tidak sesuai dengan validasi pada struct
	}

	// melakukan validasi dari struct
	err := validate.Struct(request)

	// mengecek error validasi
	if err != nil {
		fmt.Println(err.Error())
	}
}

// implementasi or rule
func TestORRule(t *testing.T) {
	// membuat object validate
	validate := validator.New()

	// membuat struct baru
	type Login struct {
		Username string `validate:"required,email|numeric"` // menggunakan operator OR
		Password string `validate:"required"`
	}

	// membuat object baru dari struct
	request := Login{
		Username: "081234567890",
		Password: "12345",
	}

	// melakukan validasi dari struct
	err := validate.Struct(request)

	// mengecek error validasi
	if err != nil {
		fmt.Println(err.Error())
	}
}

// implementasi custom validation cross field

// membuat custom function ntuk membandingkan 2 buah nilai, yang pertama dan yang kedua sama-
// tidak perludi huruf besar kecilnya
func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	// mengambil data parameter dengan method GetStructFieldOK2()
	// retun dari method GetStructFieldOK2(), berupa value, tipe data, nullable (true/false), ok (iya/tidak)
	value, _, _, ok  := field.GetStructFieldOK2()

	// mengecek input terlebih dahulu
	// kalau misalnya input nya keliru, maka akan panic
	if !ok {
		panic("field not ok")
	}

	// menyimpan data pertama dengan mengkonversinya ke interface lalu ke string
	data := field.Field().Interface().(string)

	// kemudian ambil kedua buah data, dan mengubah nya menjadi upper case (huruf besar semua)
	firstValue := strings.ToUpper(data)
	secondValue := strings.ToUpper(value.String())

	// kemudian membandingkan data pertama dan kedua
	return firstValue == secondValue
}

func TestCustomValidationCrossField(t *testing.T) {
	// membuat object validate
	validate := validator.New()

	// melakukan register validation custom dari function yang sudah dibuat sebelumnya
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	// membuat struct baru
	type User struct {
		// dengan validasi cross ini, akan membandingkan data yang kita punya dan mencocokannya, jika tidak cocok akan error validasi
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email string `validate:"required,email"`
		Phone string `validate:"required,numeric"`
		Name string `validate:"required"`
	}

	// membuat object baru dari struct
	request := User{
		Username: "081234567890",
		Email: "taufik@email.com", 
		Phone: "081234567890", 
		Name: "Taufik", 
	}

	// melakukan validasi dari struct
	err := validate.Struct(request)

	// mengecek error validasi
	if err != nil {
		fmt.Println(err.Error())
	}
}

// implementasi struct level validation

// membuat struct baru
type RegisterRequest struct {
	// memiliki attribute yang mengimplementasikan validation
	Username string `validate="required"`
	Email string `validate="required,email"`
	Phone string `validate="required,numeric"`
	Password string `validate="required"`
}

// pada saat membuat custom validation untuk struct, parameter nya tidak lagi menggunakan fieldLevel-
// akan tetapi menggunakan structlevel. 
// juga tidak perlu mewajibkan return value, sehingga peneapan ekspresi error pun berbeda
// kalau custom validation di struct jika menghasilkan error maka akan mengembalikan ReportError

// membuat custom validation untuk struct
func MustValidRegisterSuccess(level validator.StructLevel) {
	// mengambil data / value saat ini, dan mengkonversinya dalam bentuk struct
	registerRequest := level.Current().Interface().(RegisterRequest)

	// memberikan / menambahkan validasi custom, dengan memberikan pengecekan
	// jika data struct username nya sama dengan email atau username nya sama dengan phone, maka akan sukses
	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// sukses
	} else {
		// gagal, akan mengembalikan ReportError dengan parameter (value, field name, struct field name, tag(bebas), parameter(opsional))
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidatoin(t *testing.T) {
	// membuat object validate
	validate := validator.New()

	// melakukan register validation custom dari function yang sudah dibuat sebelumnya
	// khusus untuk custom validation struct menggunakan method yang berbeda dari sebelumnya
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	// membuat object baru dari struct register request
	request := RegisterRequest{
		Username: "akuutauf@email.com", // akan memunculkan error validasi, karena emailnya berbeda
		Email: "taufik@email.com", 
		Phone: "081234567890", 
		Password: "rahasia", 
	}

	// melakukan validasi dari struct
	err := validate.Struct(request)

	// mengecek error validasi
	if err != nil {
		fmt.Println(err.Error())
	}
}