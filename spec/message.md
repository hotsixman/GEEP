# Message

## Process Connect
### Connect request
```ts
type Message = {
    type: "connect"
    name: string
}
```

### Command
```ts
type Message = {
    type: "command"
    command: string
}
```

### Start
```ts
type Message = {
    type: "start"
    name: string
    run: string
    args: string[]
    cwd: string
    env: Record<string, string>
}
```