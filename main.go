package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)
func main(){
	number:=1
	output :=make([]Boss,0)
	url:=""
	for number=1;number<10;number++{
	//北京 https://www.zhipin.com/c101010100/?query=golang&page=2&ka=page-2
	//101020100 上海  101280100广州
		url = "https://www.zhipin.com/c101020100/?query=golang&page="+strconv.Itoa(number)+"&ka=page-2"
		println(url)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("cache-control", "no-cache")
		res, err := http.DefaultClient.Do(req)
		if err!=nil{
			println("GET error")
			return
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		//ioutil.WriteFile("response2.html",body,0644)
		//fmt.Println(res)
		//fmt.Println(string(body))
		/*	x:=MatchLine(string(body))
			for _,v:=range x{
				println(v)
			}
		*/

		reg := regexp.MustCompile(`job-title.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*\n.*`)//match job info
		result:=reg.FindAllString(string(body),-1)
		//println(number,result)
		for _,v:=range result{
			vv,err:=myCut(v) // clean and get job detail
			if err !=nil{
				println(err.Error(),vv.title)
			}
			output=append(output,vv)
			//myCut(v)
			//myShow(vv)
		}
	}
	println(len(output))
	Output(output)

}

func Output(result []Boss)error{
	out:=""
	for _,value:=range result{
		out=out+value.name+","+value.title+","+value.experience+","+value.salary+","+value.address+"\n"
	}

	ioutil.WriteFile("resultSH.csv",[]byte(out),0644)
	return nil;
}

//https://www.zhipin.com/job_detail/?query=golang&city=101280600&industry=&position=
//https://www.zhipin.com/c101280600/?query=golang&page=2&ka=page-2
type Boss struct{
	 title string
	salary string
	 address string
	 experience string
		name string

}
func myCut(text string )(Boss,error){
	reg,err := regexp.Compile(`>[^e\n]+<`)
	a:=Boss{}
	if err!=nil{
		return a,err
	}
	result:=reg.FindAll([]byte(text),-1)


	for i,v:=range result{
		w:=v[1:len(v)-1]
		println(i,string(w))
		if i==0{
			a.title=string(w)
		}
		if i==1{
			a.salary=string(w)
		}
		if i==2{
			a.address=string(w)
		}
		if i==3{
			a.experience=string(w)
		}
		if i==5{
			a.name=string(w)
		}

	}
	return a,nil
}