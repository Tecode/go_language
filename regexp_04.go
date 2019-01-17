package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := `<div class="flex-auto f6 mr-3">
      <a href="/google/re2/commits?author=junyer"
     class="commit-author tooltipped tooltipped-s user-mention"
     aria-label="View all commits by junyer">junyer</a>
        <a data-pjax="true" title="Avoid null PODArray&lt;&gt; issues in SparseSet and SparseArray&lt;&gt;.
Change-Id: Ibd21e4277c2855d8572762e6f6ea9b3f90159273
Reviewed-on: https://code-review.googlesource.com/c/37398
Reviewed-by: Paul Wankadia &lt;junyer@google.com&gt;" class="message" href="/google/re2/commit/68752454406a2ca8de090e41927702288dfe101e">Avoid null PODArray&lt;&gt; issues in SparseSet and SparseArray&lt;&gt;.</a>
          <span class="hidden-text-expander inline">
            <button type="button" class="ellipsis-expander js-details-target" aria-expanded="false">&hellip;</button>
          </span>
        <div class="commit-desc"><pre class="text-small">Change-Id: Ibd21e4277c2855d8572762e6f6ea9b3f90159273
Reviewed-on: <a href="https://code-review.googlesource.com/c/37398" rel="nofollow">https://code-review.googlesource.com/c/37398</a>
Reviewed-by: Paul Wankadia &lt;junyer@google.com&gt;</pre></div>
    </div>
    <div class="no-wrap">
      Latest commit
      <a class="commit-tease-sha" href="/google/re2/commit/68752454406a2ca8de090e41927702288dfe101e" data-pjax>
        6875245
      </a>
      <span itemprop="dateModified"><relative-time datetime="2019-01-11T08:41:38Z">Jan 11, 2019</relative-time></span>
    </div>
  </div>`
	reg := regexp.MustCompile(`<div class=(.*?)>`)
	match := reg.FindAllString(buf, -1)
	fmt.Println(match)


	for _, value := range match {
		fmt.Println(value)
	}

}
