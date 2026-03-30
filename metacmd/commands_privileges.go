package metacmd

func init() {
	commands = append(commands,
		command{
			Name:                "ddp",
			Usage:               "[PATTERN]",
			MinArgs:             0,
			MaxArgs:             1,
			DisablePatternRewrite: true,
			Help:                "list default access privileges",
			Run:                 execDescribeDefaultPrivileges,
		},
		command{
			Name:                "dg",
			Aliases:             []string{"du"},
			Usage:               "[PATTERN]",
			MinArgs:             0,
			MaxArgs:             1,
			Flags:               "S+",
			DisablePatternRewrite: true,
			Help:                "list roles",
			Run:                 execDescribeRoles,
		},
		command{
			Name:                "drds",
			Usage:               "[PATRN1 [PATRN2]]",
			MinArgs:             0,
			MaxArgs:             2,
			DisablePatternRewrite: true,
			Help:                "list per-database role settings",
			Run:                 execDescribeRoleSettings,
		},
	)

	commandHelp = append(commandHelp,
		helpEntry{Command: "\\ddp [PATTERN]", Description: "list default access privileges"},
		helpEntry{Command: "\\dg[S+] [PATTERN]", Description: "list roles"},
		helpEntry{Command: "\\du[S+] [PATTERN]", Description: "list roles"},
		helpEntry{Command: "\\drds [PATRN1 [PATRN2]]", Description: "list per-database role settings"},
	)
}
