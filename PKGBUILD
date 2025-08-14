pkgname=lura
pkgver=1.3
pkgrel=1
pkgdesc="Simple turn based game"
arch=('x86_64')
url="https://github.com/IwnuplyNotTyan/Lura"
license=('MIT')
depends=('ttf-nerd-fonts-symbols' 'ttf-nerd-fonts-symbols-mono')
makedepends=('go')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/IwnuplyNotTyan/Lura/archive/${pkgver}.tar.gz")
sha256sums=('SKIP')

build() {
  cd "${srcdir}/Lura-${pkgver}"
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  
  go mod download
  go build -o lura
}

package() {
  cd "${srcdir}/Lura-${pkgver}"
  install -Dm755 lura "${pkgdir}/usr/bin/lura"
  
  if [ -f LICENSE ]; then
    install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
  fi
}
