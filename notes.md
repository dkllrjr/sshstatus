- yaml file to store server name, host, and port

#

server name
- ip
- port


func main() {

	config_file := "config.yml"

	config := get_config(config_file)

	for ip, port := range config {
		fmt.Printf("%s: %s\n", ip, port)
	}

	fmt.Printf("%v\n", config["prandtl"])
}


func main() {

    host, port := os.Args[1], os.Args[2]
    sec, _ := strconv.Atoi(os.Args[3])

    fmt.Println(IsOpen(host, port, sec))
}
