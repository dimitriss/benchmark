# benchmark
Simple Go benchmark for testing


#Sample results:
#i5 2500K @ 4.0Ghz
3957809
       1        46423520800 ns/op

# RPi 2 @ 1Ghz - Debian ARMHF
3957809
       1        1102586688354 ns/op

# RPi 2 @ 1Ghz - Raspbian       
3957809
       1        1113241230692 ns/op

#explanation
3957809:prime numbers
1 46423520800 ns/op
means that the loop ran 1 time at a speed of 46423520800 ns per loop.
