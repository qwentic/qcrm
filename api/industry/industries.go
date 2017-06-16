package industry

var (
	Informationtechnology = &Industry{
		Name: "Information Technology",
	}
	Telecommunication = &Industry{
		Name: "Telecommunication",
	}
	Manufacturing = &Industry{
		Name: "Manufacturing",
	}
	BankingServices = &Industry{
		Name: "Banking Services",
	}
	Consulting = &Industry{
		Name: "Consulting",
	}
	Finance = &Industry{
		Name: "Finance",
	}
	Government = &Industry{
		Name: "Government",
	}

	Delivery = &Industry{
		Name: "Delivery",
	}
	NonProfit = &Industry{
		Name: "Non-Profit",
	}

	Entertainment = &Industry{
		Name: "Entertainment",
	}
	Other = &Industry{
		Name: "Other",
	}

	All = []*Industry{
		Informationtechnology,
		Telecommunication,
		Manufacturing,
		BankingServices,
		Consulting,
		Finance,
		Government,
		Delivery,
		NonProfit,
		Entertainment,
		Other,
	}
)
