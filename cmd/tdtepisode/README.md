# tdtepisode - Podcast publishing tools for TheDrunkenTaoist.com

## Current Workflow

Before any of the publishing can take place, these things must have already taken place:

- Episode has been recorded & edited
- Episode media has been uploaded to Wiredrive
- Daniele has written the episode notes and emailed them to me

The process of publishing a new episode goes something like: 

- Download notes and image from email
- Download media from Wiredrive
- Upload image to Wiredrive
- Upload media to CDN
- Convert the notes to plaintext
- Create episode HTML
- Publish draft of episode to site
- Email Daniele to review
- Daniele and I have back-and-forth to adjust if necessary
- Daniele approves
- Publish episode

This tool assists in the process by automating as much as possible. 

## tdtepisode Workflow

Much of the existing workflow is a side-effect of the platform being overly complex.
At the end of the day, all we want to do is take some episode notes, a thumbnail JPG/PNG
and an MP3 and produce some resources that are accessible on the internet. 

Key insight: once an episode is published, it almost never changes.

So, why not use a static generator and cut out all the complexity involved with having a Django app and DB and whatnot
and instead generate it all once, upload it to a static hosting site and be done. Git can be the database. And with it, all the
content of the site will be archived and available in "the cloud" without having to do anything other than git push.

The workflow would then become:

- run `tdtepisode gen` to generate the structure for the episode
- plug in notes, links, thumbnail, etc
- run `tdtepisode render` to render episode to html, xml, etc
- verify things look good in a local browser
- run `tdtepisode publish --preview` to upload the content to the appropriate places for review
- email Bolelli,
- adjust if necessary
- verify in a browser at the public domain
- do git add/commit/push
