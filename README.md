# rssquash

RSS/Atom/JSON feed aggregator. Consolidate multiple feeds into one easily consumable Atom feed.

## Quickstart

```
go run main.go --source <path/to/feed/list>
```

### Explanation

- `--source` Read in a newline delimited list of feed URLs

## Example

```
go run main.go --source <(printf '%s\n' "https://static.fsf.org/fsforg/rss/news.xml" "https://github.blog/all.atom")
```

## Deploy

*Note:* This is a proof of concept

```
export RSSQUASH_PREFIX=sample/ RSSQUASH_BASEURL=https://egladman.github.io/rssquash/
./deploy.sh <path/to/feed/list>
```

## Dependencies

- [gofeed](https://github.com/mmcdole/gofeed)

## License

- This project is licensed under the [MIT License](LICENSE).
  - Content in branch `gh-pages` is not licensed, and is redistributed.
