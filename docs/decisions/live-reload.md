# Live Reload

## Take 1: GoTTH Current Starter

The current iteration of the GoTTH template live-reload looks like:

- Docker Container based on golang:alpine w/ templ, air, and tailwindcss binary installed
- Docker entrypoint runs [air](https://github.com/TomDoesTech/GOTTH/) and templ watch
- Application main.go runs tailwind-build in init() function and then starts server

### Problems:

- Due to running tailwind-build before server-start, the server is down
long enough for air's proxy to run through its hard-coded 10 retries and the page
goes to an error state.
- Tailwind v4's binary is flaky in an alpine linux image. Even after setting the
`platform` in the `docker-compose.yml` file, I observed occasional "illegal
instruction" errors

## Take 2: [Templ's "Live Reload with Other Tools"](https://templ.guide/developer-tools/live-reload-with-other-tools)

Next I tried to follow Templ's guide for live reload with multiple tools, which looks like:

- Makefile with 4 targets
  - live/templ runs `templ generate --watch` w/ proxy (Inserts live-reload into HTML <body />)
  - live/server runs `air` to watch all go file changes and restart server
  - live/tailwind runs `npx tailwindcss --watch` to rebuild CSS
  - live/sync_assets runs `air` to watch the static/assets folder and notify templ proxy
- Run all 4 makefile targets w/ `make -j4 ...`

### Problems:

- The console output / log experience is awful. All 4 jobs output with no distinction
which job is writing. This was especially annoying when jobs failed and all I saw
was an exit message with no way to know which one failed.
- This approach requires node to be installed on the devs machine. This isn't a
deal breaker but I'd like it to not be a requirement when working on a go project.
- Air was FREQUENTLY failing to kill the server process and instead was just 
trying to start a new one with the old one running, which would then fail with a
`bind: address already in use` message. I've found many Github issues about this
problem spanning years and especially happening on MacOS and Windows. Each issue
says the problem was resolved at various times, so maybe this is something that 
just keeps regressing. Either way this is what really killed this approach for me.
Air just was not working properly.

## Take 3: Take What We Learned

Trying to work around all the problems we've hit so far, what I've landed on:

- Tailwind in a `node:alpine` docker container to watch/build CSS
- `templ generate --watch` in a `golang:alpine` docker container to watch/build `.templ`
- Use `--cmd` on `templ --watch` to remove need for air, and have templ proxy server
run the go server
- `docker compose up` to run everything

### How this addresses earlier problems:

- Tailwind runs async in another container so it doesn't increase server downtime
- Running `npx tailwind` in `node:alpine` is not exhibiting platform issues, but
still allows developers to work on the project without node locally installed
- Docker compose provides service-prefixed logs so you can tell what's happening
- `templ generate --watch --cmd` seems to be properly killing the server on restarts,
so we are not getting the port taken error

### Quirks

- If a developer wants to add tailwind plugins they'll need to use node to update
the `package-lock`. Up to the dev if they want to run `npm` in a docker container
or install locally.
- `tailwind` is not currently notifying the proxy of rebuilds. That's because
currently rebuilds are almost always caused by changes to `.templ` files which
are going to trigger a server restart and reload anyways, and I don't want the 
browser to refresh twice. This is working fine because tailwind build seems to
always be faster than the server restart, but if tailwind ever loses the race
or a non-templ file causes a CSS rebuild the user will have to manually refresh
the page. I'm fine with this as the alternative seems to be risking double refresh
on every change which I'd rather not do.
