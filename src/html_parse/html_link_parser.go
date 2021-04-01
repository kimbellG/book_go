package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	htm "golang.org/x/net/html"
)

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			log.Fatalln("main:", err)
		}
		defer resp.Body.Close()

		doc, err := htm.Parse(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks: %v\n", err)
			os.Exit(1)
		}

		for _, link := range visitLoopAndCheckLinks(nil, doc) {
			fmt.Println(link)
		}
	}
}

func visitLoopAndCheckLinks(links []string, n *htm.Node) []string {

	links = ScanImg(links, n)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitLoopAndCheckLinks(links, c)
	}

	return links
}

func visitRecursiveAndCheckLinks(links []string, n *htm.Node) []string {
	if n == nil {
		return links
	}

	links = ScanLinks(links, n)

	links = visitRecursiveAndCheckLinks(links, n.FirstChild)
	links = visitRecursiveAndCheckLinks(links, n.NextSibling)

	return links
}

func ScanLinks(links []string, n *htm.Node) []string {
	if isATag(n) {
		for _, a := range n.Attr {
			if isLink(a) {
				links = append(links, a.Val)
			}
		}
	}

	return links
}

func isLink(a htm.Attribute) bool {
	if a.Key == "href" {
		return !strings.HasPrefix(a.Val, "#")
	}
	return false
}

func isATag(n *htm.Node) bool {
	return n.Type == htm.ElementNode && n.Data == "a"
}

func isEqualSlice(left, right []string) bool {
	if isNotEqualSliceLen(left, right) {
		return false
	}

	for i, _ := range left {
		if isNotEqualSliceAttr(left[i], right[i]) {
			return false
		}
	}

	return true
}

func isNotEqualSliceLen(left, right []string) bool {
	return len(left) != len(right)
}

func isNotEqualSliceAttr(left, right string) bool {
	return left != right
}

func outline(stack []string, n *html.Node) {
	if isTextNode(n) {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func isTextNode(n *html.Node) bool {
	return n.Type == html.TextNode && n.Parent.Data != "script"
}

func visitRecursiveAndCheckElement(elements map[string]int, n *htm.Node) map[string]int {
	if n == nil {
		return elements
	}

	elements = CheckElement(elements, n)

	elements = visitRecursiveAndCheckElement(elements, n.FirstChild)
	elements = visitRecursiveAndCheckElement(elements, n.NextSibling)

	return elements
}

func CheckElement(elements map[string]int, n *htm.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	return elements
}

func ScanImg(linksOnImg []string, n *html.Node) []string {
	if isImg(n) {
		for _, attr := range n.Attr {
			if isLinkOnImage(&attr) {
				linksOnImg = append(linksOnImg, attr.Val)
			}
		}
	}

	return linksOnImg
}

func isImg(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
}

func isLinkOnImage(attr *html.Attribute) bool {
	return attr.Key == "src"
}
