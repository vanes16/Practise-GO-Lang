package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string) {
	for node != nil {
		fmt.Printf("%d", node.data)

		node = node.next

		if node != nil {
			fmt.Printf(sep)
		}
	}
}

/*
 * Complete the 'deleteDuplicates' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts INTEGER_SINGLY_LINKED_LIST head as parameter.
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func deleteDuplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil {
		return nil
	}

	current := head

	for current.next != nil {
		if current.data == current.next.data {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}

	return head
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	headCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	head := SinglyLinkedList{}
	for i := 0; i < int(headCount); i++ {
		headItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		headItem := int32(headItemTemp)
		head.insertNodeIntoSinglyLinkedList(headItem)
	}

	result := deleteDuplicates(head.head)

	printSinglyLinkedList(result, "\n")
	fmt.Printf("\n")
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
