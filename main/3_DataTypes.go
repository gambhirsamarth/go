package main

func main() {
	/*
			String
			In Go, a string is a sequence of bytes.

			They are declared either using double quotes
		or backticks which can span multiple lines
	*/
	var name string = "My name is sam"
	var bio string = `I am Sam.
                                    I am a software engineer.`

	/*
		Bool
		It can have two possible values - true or false.
		normal operators == for equality check != for inequality
	*/

	var value bool = true

	//SIGNED INTEGERS
	var i int = 404                     // Platform dependent
	var i8 int8 = 127                   // -128 to 127
	var i16 int16 = 32767               // -2^15 to 2^15 - 1
	var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
	var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1

	// UNSIGNED INTEGERS
	var ui uint = 404                     // Platform dependent
	var ui8 uint8 = 255                   // 0 to 255
	var ui16 uint16 = 65535               // 0 to 2^16
	var ui32 uint32 = 2147483647          // 0 to 2^32
	var ui64 uint64 = 9223372036854775807 // 0 to 2^64
	var uiptr uintptr                     // Integer representation of a memory address

	/*
		It is recommended that whenever we need an integer value,
		we should just use int unless we have a specific reason
		to use a sized or unsigned integer type.
	*/
}
