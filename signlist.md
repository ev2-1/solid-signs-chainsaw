# Documentation of signlist files

All lines starting with `#` are comments

example sign definition:
```
pos 4 10 -5@server01 wall South
text I like Döner Kebap\n%s
color green
dyn PlayerCnt:server02
end
```

here the line starting with `pos` defines a signPos

- `4 10 -5` are the coordinates of the sign
- `server01` is the server they are on
- `wall` says they are wallsigns (anything else would tell its a standing sign)
- `South` is the Direction the sign is facing, check signs mod for details

the line with `text` at the begining contains the text on the sign
- `%s` will be replaced with the dyn arguments, (when having multible dyn arguments, you use multible `%s`'es)
- `\n` will be replaced with an newline

`color`
- the color of the sign
- is a color spec string

lines starting with `dyn` contain dynamic information,
in this case the `PlayerCnt` PlayerCount of `server02`
in this case this means the %s will be replaced by the number representing the playercount on `server02`

`end` ends a sign definition and starts a new one
