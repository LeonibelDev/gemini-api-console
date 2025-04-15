/*
	Function to handle user GEMINI API KEY
	and storage
*/

package utility

import (
	"fmt"
	"os"
)

func SaveAPIKEY(){
	if os.Getenv("GEMINI_API_KEY") == ""{
		fmt.Println("User doesnt have a GEMINI API KEY")
		return
	}
}
