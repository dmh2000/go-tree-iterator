# go-tree-iterator

I was diddling with the new **range over function iterator** scheme in Golang and thought of a test case that would help me understand it. So I did some implementation and testing of a version of an iterator for a binary search tree.

Once I started on this effort, it kind of morphed into an exercise using AI to generate some or all of the code. That turned out to be really interesting.

## The Application

### Hash Map vs Tree Map

Golang has a built in hash **map** that fits most use cases for a key/value store. Iterating over a Golang **map** does not guarantee any particular order.

On the other hand, in the first cut of the C++ STL (98), the **std::map** container was implemented as a binary search tree rather than a hash map. The **std::map** had a behavior that the Go hash map does not have : iterating over the keys/value in order (https://en.cppreference.com/w/cpp/container/map). The actual implementation is usually a red-black tree. Subsequent updates of the C++ STL (11) added the **std::unordered_map** which was a hash map.

Most of the time a hash map will have insert,lookup and delete operations that O(1). That's better than a generic search tree which would have average O(log N) insert, lookup and delete, but could degenerate to O(N) if values were inserted in a particular order. A red-black tree (https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) avoids that situation by maintaining a mostly balanced tree structure. Insert, lookup and delete are worst case O(log N).

In any case, there could be specific instances where iterating over a map in order is more important than insert/delete time. For example if the structure is only built once or added infrequently, but iterating in order is done more often.

### Implementing a Red-Black Tree in Go

I figured if I was going to do this, I should bite the bullet and code up something better than a simple BST. I picked a red-black tree.

As an aside, I learned on and have always used the Sedgwick algorithms (https://algs4.cs.princeton.edu/home/) as a reference for implementation and analysis. I found the CLRS book (https://mitpress.mit.edu/9780262046305/introduction-to-algorithms/) a bit too mathy for me, maybe I'm not smart enough.

The Princeton/Sedgewick website provides a free online course, book and Java implementation for all the algorithms covered in the book and course. Highly recommended if you are learning algorithms on your own. I find it very accessible.

Since it might take me a month or two to remember how to implement a red-black tree (if I could do it at all), I did a port of the Sedgick Java implementation RBT to Go for this excercise (https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/RedBlackBST.java.html)

I also decided to make the implementation generic for key/value types.

## AI

### AI Did The Port - Disclaimer

I could have hand coded the implementation line by line, but since it's 2024, I used a bit (actually a lot) of AI to help me out. This turned out to be a better learning experience than implementing the iterator.

I tried three approaches:

#### Github Copilot - incrementally

- I have been using Github Copilot extensively and it has really accerated my development time. By a lot.
- Instead of a chat interface, I started implementing a Go version line by, modeled after the Java example.

#### Gemini Flash LLM to generate the whole thing in one swoop.

- I used the [sgpt command line tool](https://github.com/tbckr/sgpt) with the Gemini model
- I had not use this approach before. I got it from AiCodeKing https://www.youtube.com/watch?v=BoihrNkJ9dY. Btw, his series of AI videos is really great. It covers how to use all the different models and associated software. Recommended.
- Here's the prompt I used:
  - "$sgpt --code "using the file at https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/RedBlackBST.java.html as a model, create a version of that code using the go language. The implementation should be generic with respect to key and value types."

#### ChatGpt LLM to generate the whole thing in one swoop.

- I used the OpenAI ChatGPT online interface.
- I used the same prompt as I did for Gemini:
  - "$sgpt --code "using the file at https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/RedBlackBST.java.html as a model, create a version of that code using the go language. The implementation should be generic with respect to key and value types"

### AI Versions

Disclaimer : I subscribe to GitHub Copilot and Google Gemini. I used the free version for ChatGPT.

### Github Copilot

- In the first line in the file I put the link to the Java version. I also copied the Java version into a file in the same directory. I don't know if it looked a the link or the local copy or if it had already been trained on this code, because it seemd to know exactly what to gnerate.
- I began by coding line by line following the Java model verbatim. Once I had converted the 'Node' struct to a generic version and got about 30 lines into the program manually, Copilot caught on.
- After that I implemented each function by starting with a comment matching what I wanted and then Copilot did its two-step procedure. I advanced to the line after the comment and it generated the function declaration. Then on next line it suggested a solution for that function and it completed it.

### Gemini/Sgpt

- The code it generated was almost perfect. I had a very few things I needed to fix.
- It was smart enough to add error handling where I had to add it to the Github Copilot version.
- The one thing that it differed from the Copilot version that instead of returning errors (for things like 'no such key'), it return a boolean, which I translated to 'ok' in the code.
- It was clear that Gemini either used the link I gave it or it had already trained on this code because it almost word for word matched the Java version.
- Fixes
  - fix the function namesthat are marked public in the Java code so they are exported ( I could have refined my prompt to have it do that in the generator)
  - had to add **constraints.Ordered** to the key 'K' type specs. This showed up as syntax errors in the places where the keys were being compared.

### ChatGpt - free online version

- The code it generated was good but it left out a set of functions hat retrieved the list of keys.
- Like Gemini it used a bool 'ok' instead of throwing errors for things like key not found or empty tree.
- I ran two versions. For the first one I uploaded the bst.java file and prompted it to use that. This attempt did not seem to use the file because it generated a stripped down version. For the second attempt I prompted it to use the GitHub file instead of the local one and it did much better.
- Fixes
  - fix the function namesthat are marked public in the Java code so they are exported
  - since the 'Key's funcs were missing, I pulled over the Gemini versions. Did not require any changes.
  - had to add **constraints.Ordered** to the key 'K' type specs. This showed up as syntax errors in the places where the keys were being compared.

### Conclusions

Because of the Copilot line-by-line approach, that took quite a bit longer to finish. However, by doing it this way, I could see and understand what it was generating. Using the Gemini approach, I was lazy and didn't review the code other than what was needed to fix a few quirks. Same with the ChatGPT version, a bit more work to fixup than Gemini.

Once I was done, all implementations passed my minimal test cases.

This brings up the issue of TDD and confidence in an AI generated solution. Is it trustworthy without a detailed manual review? Are the tests good enough to accept an AI generated solution? Does using an AI generator or helper mean that no human will know actually what its doing or how to fix or modify it?

The time it took to code this stuff up was much faster than if I had done a manual line by line conversion. Or God forbid I had to crap out an RBT on my own.

## Implementing the Iterator

Ok, back to implementing the Range over Function iterator.

Iterating over a binary search tree is pretty simple, just do an inorder traversal and emit the key/value pairs. Something like this:

```go
func (t *xxxRBT[K, V]) Iterator() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		var inorder func(*Node[K, V])
		inorder = func(n *Node[K, V]) {
			if n == nil {
				return
			}
			inorder(n.left)
			if !yield(n.key, n.val) {
				return
			}
			inorder(n.right)
		}
		inorder(t.root)
	}
}
```

After got an iterator for each version, I realized I should go ahead and create an interface for an RBT and make them all comply.

https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/RedBlackBST.java.html
https://bitfieldconsulting.com/posts/iterators
