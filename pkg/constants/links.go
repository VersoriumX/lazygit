package constants

type Docs struct {
	CustomPagers      string
	CustomCommands    string
	CustomKeybindings string
	Keybindings       string
	Undoing           string
	Config            string
	Tutorial          string
}

var Links = struct {
	Docs        Docs
	Issues      string
	Donate      string
	Discussions string
	RepoUrl     string
	Releases    string
}{
	RepoUrl:     "https://github.com/VersoriumX/lazygit",
	Issues:      "https://github.com/VersoriumX/lazygit/issues",
	Donate:      "https://github.com/sponsors/VersoriumX",
	Discussions: "https://github.com/VersoriumX/lazygit/discussions",
	Releases:    "https://github.com/VersoriumX/lazygit/releases",
	Docs: Docs{
		CustomPagers:      "https://github.com/VersoriumX/lazygit/blob/master/docs/Custom_Pagers.md",
		CustomKeybindings: "https://github.com/VersoriumX/lazygit/blob/master/docs/keybindings/Custom_Keybindings.md",
		CustomCommands:    "https://github.com/VersoriumX/lazygit/wiki/Custom-Commands-Compendium",
		Keybindings:       "https://github.com/VersoriumX/lazygit/blob/master/docs/keybindings",
		Undoing:           "https://github.com/VersoriumX/lazygit/blob/master/docs/Undoing.md",
		Config:            "https://github.com/VersoriumX/lazygit/blob/master/docs/Config.md",
		Tutorial:          "https://www.youtube.com/watch?v=qIKyTs-vX_Q",
	},
}
