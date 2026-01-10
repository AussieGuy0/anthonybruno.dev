# Strava Integration

Fetches activities from Strava and generates Jekyll collection files.

## Setup

1. **Get Strava API credentials** at https://www.strava.com/settings/api

2. **Authorize the app** (first time only):
   ```bash
   cd _scripts/strava/authorize
   export STRAVA_CLIENT_ID=your_client_id
   export STRAVA_CLIENT_SECRET=your_client_secret
   go run authorize.go
   ```
   Follow the prompts to authorize and save credentials to `.env`

3. **Fetch activities**:
   ```bash
   make fetch_strava
   ```

## Environment Variables

Create `.env` with:
```
STRAVA_CLIENT_ID=your_client_id
STRAVA_CLIENT_SECRET=your_client_secret
STRAVA_REFRESH_TOKEN=your_refresh_token
```

## GitHub Actions

Add these secrets to your repository for automated syncing:
- `STRAVA_CLIENT_ID`
- `STRAVA_CLIENT_SECRET`
- `STRAVA_REFRESH_TOKEN`

