from collections import deque

def mergeSort(array):
    if len(array) <= 1:
        return array
    
    midIdx = len(array) // 2
    leftSorted = mergeSort(array[:midIdx])
    rightSorted = mergeSort(array[midIdx:])
    return merge(leftSorted, rightSorted)

# merge 2 sorted array to a single sorted array
def merge(arrayOne, arrayTwo):
    arrayOne = deque(arrayOne)
    arrayTwo = deque(arrayTwo)

    merged = []
    while arrayOne and arrayTwo:
        # remove the front element 
        if arrayOne[0] < arrayTwo[0]:
            merged.append(arrayOne.popleft())
        else:
            merged.append(arrayTwo.popleft())
    
    merged += arrayOne
    merged += arrayTwo

    return merged


if __name__ == "__main__":
    array = [9, 5, 2, 4, 11, 8, 22]
    print(mergeSort(array))

"""
output: [2, 4, 5, 8, 9, 11, 22]
"""