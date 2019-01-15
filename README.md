# kodemo
This branch holds the second phase of the demo, which covers ko apply to deploy and watch.

## Script


```
# Show the changes to the app
cat ./cmd/myapp/main.go

# Show our knative service configuration
cat ./config/service.yaml

# Apply the yaml
time ko apply -f config

# Open it in a browser


# Start the watch
ko apply -W -f config

# Edit the config to change the environment variable.
# No build, just reapplied.
# Refresh browser tab.

# Edit main.go to add punctuation.
# Build, push, redeploy.
# Refresh browser tab.

```
