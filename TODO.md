# TODO

## General
> General options
- [x] Proper logger
- [x] Flag that sets the debug mode
- [x] Flag to show the version
- [x] Page versioning in a really basic way
- [x] Version listing
- [x] Makefile helper
- [ ] CI (gofmt, golint, more)?

## Bootstrap
> Basic elements
- [x] Echo server
- [x] Sample Model
- [x] Store interface (gorm, file...)
    - [x] Open db store
- [x] Templates

## Must
> Must have things (as tasks)
- [ ] List history (in every detail page, a link to the history should be included)
- [ ] Deal with static content (images, attachments...)
- [ ] If you're about to save not the latest version, button: save -> restore
- [ ] Delete any page or version
- [ ] Not allow to create a kind of forbidden-words list (edit, list, version...)
- [ ] Redirect page/ -> page

## Enhacements
> Things that would be awesome to have, but not necessary for now
- [ ] Config file (to select store and some other future options)
- [ ] Flag to select store
    - [ ] Open file store
- [ ] Reload if files changes [idea](https://medium.com/@olebedev/live-code-reloading-for-golang-web-projects-in-19-lines-8b2e8777b1ea#.gok9azrg4)
- [ ] Add a `make release` Makefile command
  - [idea](https://github.com/miekg/dns/blob/master/Makefile.fuzz)
  - [idea](https://github.com/miekg/dns/blob/master/Makefile.release)
