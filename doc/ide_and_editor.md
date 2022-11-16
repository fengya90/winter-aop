# How to configure IDE and editors

No need to do anything if you don't care about your `around function`.

If you want to debug your `around` function, then you need to add some configuration.

## GoLand

1. open `Run -> Edit Configurations -> Go Build -> your configuration`
2. add `-tags=winter_aop` to `Go tool arguments`
3. open `Preferences -> Go -> Build Tags & Vendoring` 
4. add `winter_aop` to `Custom tags`

## Visual Studio Code

Create/Update the configuration in the `.vscode` .

.vscode/launch.json

```
{
    "configurations": [
        {
            ...
            "buildFlags": "-tags=winter_aop"

        }
    ]
}
```

.vscode/settings.json

```
{
    "go.buildFlags": [
        "-tags=winter_aop"
    ]
}
```
