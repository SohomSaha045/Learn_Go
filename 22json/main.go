package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int	
	Platform string	`json:"website"`
	Password string	`json:"-"`
	Tags []string	`json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json")
	// encodingJson()
	decodingJson()
}

//encoding json

func encodingJson()  {
	courses := []course{
		{
			"ReactJS",
			399,
			"Udemy",
			"abc123",
			[]string{"web-dev","js"},
		},
		{
			"MERN",
			999,
			"Udemy",
			"abc123",
			[]string{"web-dev","js","fullstack"},
		},
		{
			"React-Native",
			599,
			"Udemy",
			"abc123",
			[]string{"mobile-dev","js"},
		},
		{
			"Sleanium",
			599,
			"Udemy",
			"abc123",
			nil,
		},
	}
	// fmt.Println("Courses:",courses)
	finalJson,err := json.MarshalIndent(courses,"","\t")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)	
}

func decodingJson(){
	jsonData1 := []byte(`
{	
   
                "coursename": "React-Native",
                "Price": 599,
                "website": "Udemy",
                "tags": [
                        "mobile-dev",
                        "js"
                ]
	
}
	`)
	jsonData := []byte(`
	[
	{
                "coursename": "React-Native",
                "Price": 599,
                "website": "Udemy",
                "tags": [
                        "mobile-dev",
                        "js"
                ]
        },
        {
                "coursename": "Sleanium",
                "Price": 599,
                "website": "Udemy"
    }
	]
	`)

	var courses []course 

	res:=json.Valid(jsonData)
	if res {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &courses)
		fmt.Println(courses)
		fmt.Printf("%#v\n",courses)
	}else{
		fmt.Println("Something went wrong")

	}

	// some cases where u just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData1,&myOnlineData)
	fmt.Printf("\n \n %#v \n",myOnlineData)
   
	for key,val := range myOnlineData {
		fmt.Println("Key :", key, "\tValue:", val)
	}

}