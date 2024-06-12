if err := db.Delete("users", ""); err != nil {
		fmt.Println("Error", err)
	}