module github.com/setanarut/v

go 1.24.0

retract (
	[v1.0.0, v1.1.0] // NegX() and AbsY() returns wrong axis.
)