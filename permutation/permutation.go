package permutation

/*
https://en.wikipedia.org/wiki/Heap%27s_algorithm
procedure generate(k : integer, A : array of any):
if k = 1 then
	output(A)
else
	// Generate permutations with kth unaltered
	// Initially k == length(A)
	generate(k - 1, A)

	// Generate permutations for kth swapped with each k-1 initial
	for i := 0; i < k-1; i += 1 do
		// Swap choice dependent on parity of k (even or odd)
		if k is even then
			swap(A[i], A[k-1]) // zero-indexed, the kth is at k-1
		else
			swap(A[0], A[k-1])
		end if
		generate(k - 1, A)

	end for
end if
One can also write the algorithm in a non-recursive format.[3]

procedure generate(n : integer, A : array of any):
//c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k+1, A) is called
c : array of int

for i := 0; i < n; i += 1 do
	c[i] := 0
end for

output(A)

//i acts similarly to the stack pointer
i := 0;
while i < n do
	if  c[i] < i then
		if i is even then
			swap(A[0], A[i])
		else
			swap(A[c[i]], A[i])
		end if
		output(A)
		//Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
		c[i] += 1
		//Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
		i := 0
	else
		//Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
		c[i] := 0
		i += 1
	end if
end while
*/
