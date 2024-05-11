K- tape Turing Machine Simulation in GoLang

to run clone the repository

write down your TM instructions in instructions.txt with the following rules
!!  z0,[t1,t2,t3...tk],z',[nt1,nt2,nt3..ntk],[I1,I2,I3,..Ik] !!
z0 -> denotes the current state
[t1,t2..] -> denotes the heads of the tapes which TM is looking at currently
z' -> next state 
[nt1,nt2,..ntk] -> denotes what should be written on each tape
[I1,I2,I3..Ik..] -> denotes the instructions L,R,N

sample instruction looks like this
z0,[1,#],q0,[#,1],[R,R]
z0,[#,#],halt,[#,#],[N,N]

sample input 2 input tape looks like this
111
###

this will append the contents of the first tape to the second tape and will delete the first tape

to run the program simply use go run .