# loading cue code

| pkg.func              | inputs                  | output          |
| :-------------------- | :---------------------- | :-------------- |
| cuecontext.New()      | -                       | cue.Context     |
| load.Instances()      | []string, Configuration | build.Instances |
| ctx.BuildInstance(bi) | BuildInstance           | cue.Value       |
| value.Validate()      | -                       | -               |
|                       |                         |                 |
|                       |                         |                 |

Attempt to load `album.json`: no error but `value: _`

Looks like .json file ended up orphaned files.
