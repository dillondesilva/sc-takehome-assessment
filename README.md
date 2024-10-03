# sc-interns-2024

This is my submission for the SafetyCulture 2024 Internship Application. Thanks for stopping by and please feel free to probe this README to better understand my approach towards implementing both components.

## Quick note on an assumption I made

If either the source or destination folder given to `MoveFolder` exists in more than two organisations, an error is returned.

## Quick discussion on my implementation

For Component 1, I decided to implement `GetAllChildFolders()` using a BFS style algorithm to handle some of the edge cases I developed. One of the testcases I created consisted of a "broken" folder structure which was missing some folders to keep the tree connected (see `TestGetAllChildFoldersWithSkippedFoldersInPath` in `folder/get_folder_test.go`). Following this, I initially thought that checking the prefix of each path to obtain child paths would suffice however, this would've failed this test and included folders which were not reachable. Using a standard graph traversal algorithm allows me to guarantee that any child folders can be reached and that's why I opted in for this approach.

This first implementation of Component 1 has its time complexity upper-bounded by O($n^2$) as I am checking what the adjacent subfolders to a given folder are in every iteration. This can be improved by doing some preprocessing before the traversal and constructing an adjacency list which would give a better bounded runtime of O($n+m$) (where $n$ is number of tree nodes and $m$ is number of edges in tree)

To develop Component 2, I use conditional logic to validate that it is safe to move folder contents and then utilise my `GetAllChildFolders` implementation to obtain contents of the source folder. From here, moving folders is somewhat more simple as I'm just modifying the path for each of these child folders (and the source folder) and then appending it into a slice to store the new folder structure. Following this, I then just add any unhandled folders into the new structure without modification. This is also running in O($n^2)
