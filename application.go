package main

// type Uuid struct {
//   value string
// }

type Application struct {
	UUID      string `json:"uuid,omitempty"`
	Title     string `json:"title,omitempty"`
	Type      string `json:"type,omitempty"`
	Permalink string `json:"permalink,omitempty"`
}

type Applications struct {
	Entities []Application `json:"entities,omitempty"`
	Order    []string      `json:"order,omitempty"`
}

var apps Applications = Applications{
  Entities: []Application{
    Application{
      UUID: "9a30119-d673-4978-b393-f608fe28ae07",
      Title: "Home",
      Type: "system",
      Permalink: "/",
    },
    Application{
      UUID: "54789c07-bb43-4db4-8b2d-1a8e1f8c67f1",
      Title: "Users",
      Type: "system",
      Permalink: "/apps/users",
    },
  },
  Order: []string{"9a30119-d673-4978-b393-f608fe28ae07", "54789c07-bb43-4db4-8b2d-1a8e1f8c67f1"},
}
