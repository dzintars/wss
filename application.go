package main

// type Uuid struct {
//   value string
// }

// Application ...
type Application struct {
	UUID      string `json:"uuid,omitempty"`
	Title     string `json:"title,omitempty"`
	Component string `json:"component,omitempty"`
	Permalink string `json:"permalink,omitempty"`
}

// Applications ...
type Applications struct {
	Entities map[string]Application `json:"entities,omitempty"`
	IDs      []string               `json:"ids,omitempty"`
}

var apps Applications = Applications{
	Entities: map[string]Application{
		"9a30119-d673-4978-b393-f608fe28ae07": {
			UUID:      "9a30119-d673-4978-b393-f608fe28ae07",
			Title:     "Home",
			Component: "app-home",
			Permalink: "/",
		},
		"54789c07-bb43-4db4-8b2d-1a8e1f8c67f1": {
			UUID:      "54789c07-bb43-4db4-8b2d-1a8e1f8c67f1",
			Title:     "Dispatch",
			Component: "app-users",
			Permalink: "/apps/users",
		},
		"c178025e-a209-4c50-8c34-36d35f36494c": {
			UUID:      "c178025e-a209-4c50-8c34-36d35f36494c",
			Title:     "Sales",
			Component: "app-users",
			Permalink: "/apps/users/1",
		},
		"437642dd-7d74-4213-af76-b51fc24eff0": {
			UUID:      "437642dd-7d74-4213-af76-b51fc24eff0",
			Title:     "Accounting",
			Component: "app-users",
			Permalink: "/apps/users/2",
		},
		"5a2192a0-0051-46a1-85e7-17245ba24f55": {
			UUID:      "5a2192a0-0051-46a1-85e7-17245ba24f55",
			Title:     "Settings",
			Component: "app-signin",
			Permalink: "/signin",
		},
	},
	IDs: []string{
		"9a30119-d673-4978-b393-f608fe28ae07",
		"54789c07-bb43-4db4-8b2d-1a8e1f8c67f1",
		"c178025e-a209-4c50-8c34-36d35f36494c",
		"437642dd-7d74-4213-af76-b51fc24eff0",
		"5a2192a0-0051-46a1-85e7-17245ba24f55",
	},
}
