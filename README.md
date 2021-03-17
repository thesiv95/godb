# If you run this from VSCode

1. Add folder to workspace
2. Set in *settings.json*:

```
"go.useLanguageServer": false,
    "gopls": {
        "experimentalWorkspaceModule": true
    }
```

# Warning
File .env with credentials to remote PostgreSQL DB is not included in this repository due to security reasons

Remote DB: https://www.elephantsql.com