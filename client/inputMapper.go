package client

// Activity holds the data for discord rich presence
type Activity struct {
	// What the player is currently doing
	Details string
	// The user's current party status
	State string
	// The id for a large asset of the activity, usually a snowflake
	LargeImage string
	// Text displayed when hovering over the large image of the activity
	LargeText string
	// The id for a small asset of the activity, usually a snowflake
	SmallImage string
	// Text displayed when hovering over the small image of the activity
	SmallText string
	// Information for the current party of the player
	Party *Party
	// Unix timestamps for start and/or end of the game
	Timestamps *Timestamps
	// Secrets for Rich Presence joining and spectating
	Secrets *Secrets
}

// Party holds information for the current party of the player
type Party struct {
	// The ID of the party
	ID string
	// Used to show the party's current size
	Players int
	// Used to show the party's maximum size
	MaxPlayers int
}

// Timestamps holds unix timestamps for start and/or end of the game
type Timestamps struct {
	// unix time (in milliseconds) of when the activity started
	Start int64
	// unix time (in milliseconds) of when the activity ends
	End int64
}

// Secrets holds secrets for Rich Presence joining and spectating
type Secrets struct {
	// The secret for a specific instanced match
	Match string
	// The secret for joining a party
	Join string
	// The secret for spectating a game
	Spectate string
}

func mapActivity(activity *Activity) *PayloadActivity {
	if activity.LargeImage == "" {
		activity.LargeImage = "none"
	}
	if activity.LargeText == "" {
		activity.LargeText = "none"
	}
	if activity.SmallImage == "" {
		activity.SmallImage = "none"
	}
	if activity.SmallText == "" {
		activity.SmallText = "none"
	}

	final := &PayloadActivity{
		Details: activity.Details,
		State:   activity.State,
		Assets: PayloadAssets{
			LargeImage: activity.LargeImage,
			LargeText:  activity.LargeText,
			SmallImage: activity.SmallImage,
			SmallText:  activity.SmallText,
		},
	}

	if activity.Timestamps != nil {
		final.Timestamps = &PayloadTimestamps{
			Start: uint64(activity.Timestamps.Start),
			End:   uint64(activity.Timestamps.End),
		}
	}

	if activity.Party != nil {
		final.Party = &PayloadParty{
			ID:   activity.Party.ID,
			Size: [2]int{activity.Party.Players, activity.Party.MaxPlayers},
		}
	}

	if activity.Secrets != nil {
		final.Secrets = &PayloadSecrets{
			Join:     activity.Secrets.Join,
			Match:    activity.Secrets.Match,
			Spectate: activity.Secrets.Spectate,
		}
	}

	return final
}
