package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if A > G || C < E || D < F || B > H {
		return (C-A)*(D-B) + (G-E)*(H-F)
	}
	return (C-A)*(D-B) + (G-E)*(H-F) - (min(C, G)-max(A, E))*(min(D, H)-max(B, F))
}
