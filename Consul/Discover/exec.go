package Discover

import "fmt"

func Exec(cli *client) error{
    if err := cli.Register([]string{"Go","Awesome"});err != nil{
		return err
	}

	entries,_,err := cli.Service("discovery","Go")
    fmt.Println(entries)

	if err != nil{
       return err
	}

	for _, entry := range entries{
         fmt.Println("%v\n",entry.Service)
	}
	return nil

}
