/**
 Do not return anything, modify nums1 in-place instead.
 */
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  let last = m + n - 1;

  while (m > 0 && n > 0) {
    if (nums1[m - 1] >= nums2[n - 1]) {
      nums1[last] = nums1[m - 1];
      m--;
    } else {
      nums1[last] = nums2[n - 1];
      n--;
    }
    last--;
  }

  // Fill nums1 with the remaining nums2
  while (n > 0) {
    nums1[last] = nums2[n - 1];
    n--;
    last--;
  }
}

console.log(merge([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3));
