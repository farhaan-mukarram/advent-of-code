package main

import (
	"slices"
	"strconv"
)

type FreeSpaceData struct {
	start int
	end   int
}

func findFreeSpace(blockArr []string, endIndex int, minFreeSpaceLength int) (int, int) {
	var freeSpaceDetails []FreeSpaceData
	freeSpaceStartIdx := -1
	freeSpaceEndIdx := -1

	buf := ""

	for i := 0; i < endIndex; i++ {
		ch := string(blockArr[i])

		if !areStringsEqual(ch, ".") {
			if len(buf) > 0 {
				// Save free space details
				freeSpaceEndIdx = i - 1
				freeSpaceInfo := FreeSpaceData{start: freeSpaceStartIdx, end: freeSpaceEndIdx}
				freeSpaceDetails = append(freeSpaceDetails, freeSpaceInfo)

				// clear buffer
				buf = ""
			}
		} else {
			// Save start index if first item in buffer
			if len(buf) == 0 {
				freeSpaceStartIdx = i
			}

			buf += ch
		}

	}

	// Found at least one free space slot
	if len(freeSpaceDetails) > 0 {
		freeSpaceSlotInfo := FreeSpaceData{start: -1, end: -1}

		// Find max free space slot
		for i := 0; i < len(freeSpaceDetails); i++ {
			freeSpaceItem := freeSpaceDetails[i]
			delta := freeSpaceItem.end - freeSpaceItem.start + 1

			// Found space
			if delta >= minFreeSpaceLength {
				return freeSpaceItem.start, freeSpaceItem.end
			}
		}

		return freeSpaceSlotInfo.start, freeSpaceSlotInfo.end

	}

	return freeSpaceStartIdx, freeSpaceEndIdx
}

func part2(diskMap string) int {
	// Map to store file id along with boolean that represents whether file has been processed or not
	fileIdsMap := make(map[int]bool)
	var fileIds []int

	blockArray := convertDiskMapToBlockArray(diskMap)

	// populate file ids map
	for _, c := range blockArray {
		ch := string(c)

		if areStringsEqual(ch, ".") {
			continue
		}

		fileId := toInt(ch)

		_, exists := fileIdsMap[fileId]

		if !exists {
			fileIdsMap[fileId] = false
			fileIds = append(fileIds, fileId)
		}

	}

	// Iterate over file ids in descending order
	for _, fileId := range slices.Backward(fileIds) {
		buffer := ""
		var numStartIdx, numEndIdx int

		for i := len(blockArray) - 1; i >= 0; i-- {
			c := blockArray[i]
			ch := string(c)

			if areStringsEqual(ch, ".") {
				// clear buffer and save start index
				if len(buffer) > 0 {
					numStartIdx = i + 1
					buffer = ""
				}
				continue
			}

			num := toInt(ch)

			if num == fileId {
				// empty buffer, save ending index
				if len(buffer) == 0 {
					numEndIdx = i
				}
				buffer += strconv.Itoa(num)

			} else {
				// clear buffer and save start index
				if len(buffer) > 0 {
					numStartIdx = i + 1
					buffer = ""
				}
			}

		}

		numOfFileBlocks := numEndIdx + 1 - numStartIdx
		freeSpaceStart, freeSpaceEnd := findFreeSpace(blockArray, numEndIdx, numOfFileBlocks)

		if freeSpaceEnd == -1 || freeSpaceStart == -1 {
			continue
		}

		amountOfFreeSpace := freeSpaceEnd + 1 - freeSpaceStart

		// copy num into free space if space available
		if amountOfFreeSpace >= numOfFileBlocks {
			for i := 0; i < numOfFileBlocks; i++ {
				blockArray[freeSpaceStart+i], blockArray[numStartIdx+i] = blockArray[numStartIdx+i], blockArray[freeSpaceStart+i]
			}
		}

	}

	fileChecksum := getFileChecksum(blockArray)
	return fileChecksum
}
