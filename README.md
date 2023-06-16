# quick-sort
Sorting is the process of organizing elements in a structured manner. Quicksort is one of the most popular sorting algorithms that uses nlogn comparisons to sort an array of n elements in a typical situation. Quicksort is based on the divide-and-conquer strategy. We’ll take a look at the Quicksort algorithm in this tutorial and see how it works.
this code is for learning purpose and implemented by Go language
QuickSort is a sorting algorithm based on the Divide and Conquer algorithm that picks an element as a pivot and partitions the given array around the picked pivot by placing the pivot in its correct position in the sorted array.
How does QuickSort work?
The key process in quickSort is a partition(). The target of partitions is to place the pivot (any element can be chosen to be a pivot) at its correct position in the sorted array and put all smaller elements to the left of the pivot, and all greater elements to the right of the pivot.

This partition is done recursively which finally sorts the array. See the below image for a better understanding.

Benefits of Quicksort
Let’s go through a few key benefits of using Quicksort:

It works rapidly and effectively.
It has the best time complexity when compared to other sorting algorithms.
Quick sort has a space complexity of O(logn), making it an excellent choice for situations when space is limited.
