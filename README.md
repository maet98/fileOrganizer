## File Organizer

### Installation:

Run script.sh .

```bash
chmod +x script.sh
./script.sh
```

After that, enable and start the deamon.

```bash
systemctl --user enable fileOrganizer.service
systemctl --user start fileOrganizer.service
```


### Configuration
The config file is config.json. The first key **"default"** is the directory where the deamon would be listening and moving the files and directories to their destination.

To add a new type of file you have to put the file extencion as key and the directory where you want to put it as value.
```json
{
    "pdf": "/Documents/pdf"
}
```

**Note:**
It is not neccesary to put the "~" to said that is the home directory.
