# Alien Invasion simulation

Program simulates an alien invasion of a world.

## Usage
```
// install the command
go install ./cmd/maliens

// run the command with example map
maliens ./assets/map.txt --aliens 3
```

## Expected output
```
City Foo has been destroyed by aliens alien3, alien1
---
Invasion failed - cities still exist after 10000 iterations, following is the map of the remains
---
Bar east=Baj south=Dap
Baj east=Cal west=Bar south=Pad
Cal west=Baj
Kal
Car
Baz east=Dap
Dap north=Bar west=Baz
Pad north=Baj

---
```