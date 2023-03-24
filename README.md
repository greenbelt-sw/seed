# Development Seeding
This is a app for seeding the local development server, currently with users, companies, and returns.

## Instructions
1. Pull this repo
2. Add an `.env` with authorized `MONGO_URI`
3. Run seeding tasks with `go build main.go`

### More Options
- Modify `constants.go` - specifically, counts - to update seeding amounts.
- Modify `generate.go` to update how specific data generation.