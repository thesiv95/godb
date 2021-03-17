# If you run this from VSCode

1. Add root folder of project to workspace (you should not change default workspace name and settings)
2. Set in *settings.json*:

```
"go.useLanguageServer": false,
    "gopls": {
        "experimentalWorkspaceModule": true
    }
```
3. Client part runs separately via [*LiveServer*](vscode:extension/ritwickdey.LiveServer) plugin

```
http://127.0.0.1:5500/client
```

# Warning
File `.env` with credentials to [remote PostgreSQL DB](https://www.elephantsql.com) is not included in this repository due to security reasons.
However, you can set your own DB configuration and create table **users**:
```
create table users(userid serial, name varchar, lastname varchar, birthdate varchar)
```