package scheduler

import "fmt"

type job struct {
	name         string
	cronSchedule string
	cronFunc     func()
}

func getJobs() *[]job {
	return &[]job{
		job{"Get google data", "@every 1s", getGoogleData},
		job{"SLACK", "@every 1s", getSlackData},
		job{"LINKEDIN", "@every 3s", getLinkedinData},
		job{"FACEBOOK", "@every 4s", getFacebookData},
		job{"INSTAGRAM", "@every 5s", getInstaData},
	}
}

func getGoogleData() {
	fmt.Println("GET GOOGLE DATA")
}

func getSlackData() {
	fmt.Println("GET SLACK DATA")
}

func getLinkedinData() {
	fmt.Println("GET LINKEDIN DATA")
}

func getFacebookData() {
	fmt.Println("GET Facebook DATA")
}

func getInstaData() {
	fmt.Println("GET Instagram DATA")
}
