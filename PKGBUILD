pkgname=Lura
pkgver=1.0.0
pkgrel=1
pkgdesc="Simple turn based game ~"
arch=('x86_64')
url="https://github.com/iwnuplynottyan/lura"
license=('MIT')
depends=()
makedepends=('go' 'git')
source=("git+https://github.com/iwnuplynottyan/lura.git")
sha256sums=('SKIP')

build() {
  cd "$srcdir/$pkgname"
  go mod tidy
  go build -o lura
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 lura "$pkgdir/usr/bin/lura"
}
