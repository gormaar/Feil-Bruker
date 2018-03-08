package main                                                                              
                                                                                          
import (                                                                                  
	"fmt"                                                                                 
	"os"                                                                                  
	"log"                                                                                 
	"io/ioutil"                                                                           
	"strings"                                                                             
	"strconv"                                                                             
)                                                                                         
                                                                                          
func main() {                                                                             
	file, err := os.OpenFile("notes.txt", os.O_APPEND| os.O_RDWR, 0644)                   
	if err != nil {                                                                       
		log.Fatal(err)                                                                    
	}                                                                                     
  	b,_ := ioutil.ReadAll(file)                                                    
	f:= string(b)                                                                  
	s := strings.Split(f,"\n")                                                     
	input1, input2 := s[0],s[1]                                                    
//fmt.Println(input1,input2)                                                     
                                                                                   
	tall1, _ := strconv.Atoi(input1)                                               
	tall2, _ := strconv.Atoi(input2)                                               
                                                                                   
	sum:= tall1+tall2                                                              
  fmt.Println(sum)                                                               
                                                                                                                                                                    
	err1 := os.Truncate("notes.txt", 0)                                            
	if err1 != nil {                                                               
		log.Fatal(err1)                                                            
	}                                                                              
	if _,err :=  file.WriteString(fmt.Sprintf("%d",sum)); err!= nil {              
		log.Fatal(err)                                                             
	}                                                                                                                                                               
	defer file.Close()
}
