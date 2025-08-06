package main

import "mingdemo/framework"

func SubjectAddController(c *framework.Context) error {
	c.SetOkStatus().Json("Ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.SetOkStatus().Json("Ok, SubjectListController")
	return nil
}

func SubjectDeleteController(c *framework.Context) error {
	c.SetOkStatus().Json("Ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetOkStatus().Json("Ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.SetOkStatus().Json("Ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetOkStatus().Json("Ok, SubjectNameController")
	return nil
}
