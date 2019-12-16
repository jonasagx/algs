#include <stdio.h>
#include <stdlib.h>

void swap(int v[], int i, int j) {
	int temp;

	temp = v[i];
	v[i] = v[j];
	v[j] = temp;
}

void printList(int v[], int n) {
	for(int i = 0; i < n; i++) {
		printf("%d ", v[i]);
	} 
	printf("\n");	
}

void quicksort(int v[], int n) {
	printList(v, n);

	int i, last;

	if(n <= 1) {
		return;
	}

	swap(v, 0, rand()%n);

	last = 0;
	for(i = 1; i < n; i++){
		if(v[i] < v[0]) {
			swap(v, ++last, i);
		}
	}
	swap(v, 0, last);
	quicksort(v, last);
	quicksort(v+last+1, n-last-1);
}

int isSorted(int v[], int n) {
	for(int i = 0; i < n-1; i++){
		if(v[i] > v[i+1]) {
			return 0;
		}
	}

	return 1;
}

int main(){
	int size = 5;

	int tab[5] = {5, 2, 3, 1, 4};
	quicksort(tab, size);
	printList(tab+2, size);
	printf("is sorted: %d\n", isSorted(tab, size));

	return 0;
}