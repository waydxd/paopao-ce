<template>
  <header class="header-section d-none d-xl-block">
    <div class="header-wrapper">
      <div
        class="header-bottom header-bottom-color--golden section-fluid sticky-header sticky-color--golden"
      >
        <div class="container-fluid">
          <div class="row">
            <div
              class="col-12 d-flex align-items-center justify-content-between"
            >
              <!-- Start 头部图标 -->
              <div class="header-logo">
                <div class="logo">
                  <a href="/" @click.prevent="goLink('/')">
                    <div class="svg-background"></div>
                  </a>
                </div>
              </div>
              <!-- End 头部图标 -->

              <!-- Start 头部菜单 -->
              <div
                v-show="ismenuBar"
                class="main-menu menu-color--black menu-hover-color--golden"
              >
                <nav>
                  <ul class="top-bar">
                    <li v-if="false" class="has-dropdown">
                      <a class="active main-menu-link" href="index.html"
                        >Home <i class="fa fa-angle-down"></i
                      ></a>
                      <!-- Sub Menu -->
                      <ul class="sub-menu">
                        <li><a href="index.html">Home 1</a></li>
                        <li><a href="index-2.html">Home 2</a></li>
                        <li><a href="index-3.html">Home 3</a></li>
                        <li><a href="index-4.html">Home 4</a></li>
                      </ul>
                    </li>
                    <li v-if="false" class="has-dropdown has-megaitem">
                      <a href="product-details-default.html"
                        >Shop <i class="fa fa-angle-down"></i
                      ></a>
                      <!-- Mega Menu -->
                      <div class="mega-menu">
                        <ul class="mega-menu-inner">
                          <!-- Mega Menu Sub Link -->
                          <li class="mega-menu-item">
                            <a href="#" class="mega-menu-item-title"
                              >Shop Layouts</a
                            >
                            <ul class="mega-menu-sub">
                              <li>
                                <a href="shop-grid-sidebar-left.html"
                                  >Grid Left Sidebar</a
                                >
                              </li>
                              <li>
                                <a href="shop-grid-sidebar-right.html"
                                  >Grid Right Sidebar</a
                                >
                              </li>
                              <li>
                                <a href="shop-full-width.html">Full Width</a>
                              </li>
                              <li>
                                <a href="shop-list-sidebar-left.html"
                                  >List Left Sidebar</a
                                >
                              </li>
                              <li>
                                <a href="shop-list-sidebar-right.html"
                                  >List Right Sidebar</a
                                >
                              </li>
                            </ul>
                          </li>
                          <!-- Mega Menu Sub Link -->
                          <li class="mega-menu-item">
                            <a href="#" class="mega-menu-item-title"
                              >Other Pages</a
                            >
                            <ul class="mega-menu-sub">
                              <li><a href="cart.html">Cart</a></li>
                              <li><a href="empty-cart.html">Empty Cart</a></li>
                              <li><a href="wishlist.html">Wishlist</a></li>
                              <li><a href="compare.html">Compare</a></li>
                              <li><a href="checkout.html">Checkout</a></li>
                              <li><a href="login.html">Login</a></li>
                              <li><a href="my-account.html">My Account</a></li>
                            </ul>
                          </li>
                          <!-- Mega Menu Sub Link -->
                          <li class="mega-menu-item">
                            <a href="#" class="mega-menu-item-title"
                              >Product Types</a
                            >
                            <ul class="mega-menu-sub">
                              <li>
                                <a href="product-details-default.html"
                                  >Product Default</a
                                >
                              </li>
                              <li>
                                <a href="product-details-variable.html"
                                  >Product Variable</a
                                >
                              </li>
                              <li>
                                <a href="product-details-affiliate.html"
                                  >Product Referral</a
                                >
                              </li>
                              <li>
                                <a href="product-details-group.html"
                                  >Product Group</a
                                >
                              </li>
                              <li>
                                <a href="product-details-single-slide.html"
                                  >Product Slider</a
                                >
                              </li>
                            </ul>
                          </li>
                          <!-- Mega Menu Sub Link -->
                          <li class="mega-menu-item">
                            <a href="#" class="mega-menu-item-title"
                              >Product Types</a
                            >
                            <ul class="mega-menu-sub">
                              <li>
                                <a href="product-details-tab-left.html"
                                  >Product Tab Left</a
                                >
                              </li>
                              <li>
                                <a href="product-details-tab-right.html"
                                  >Product Tab Right</a
                                >
                              </li>
                              <li>
                                <a href="product-details-gallery-left.html"
                                  >Product Gallery Left</a
                                >
                              </li>
                              <li>
                                <a href="product-details-gallery-right.html"
                                  >Product Gallery Right</a
                                >
                              </li>
                              <li>
                                <a href="product-details-sticky-left.html"
                                  >Product Sticky Left</a
                                >
                              </li>
                              <li>
                                <a href="product-details-sticky-right.html"
                                  >Product Sticky right</a
                                >
                              </li>
                            </ul>
                          </li>
                        </ul>
                        <div class="menu-banner">
                          <a href="" class="menu-banner-link">
                            <img
                              class="menu-banner-img"
                              src="assets/images/banner/menu-banner.jpg"
                              alt=""
                            />
                          </a>
                        </div>
                      </div>
                    </li>
                    <li v-if="false" class="has-dropdown">
                      <a href="blog-single-sidebar-left.html"
                        >Blog <i class="fa fa-angle-down"></i
                      ></a>
                      <!-- Sub Menu -->
                      <ul class="sub-menu">
                        <li>
                          <a href="blog-grid-sidebar-left.html"
                            >Blog Grid Sidebar left</a
                          >
                        </li>
                        <li>
                          <a href="blog-grid-sidebar-right.html"
                            >Blog Grid Sidebar Right</a
                          >
                        </li>
                        <li>
                          <a href="blog-full-width.html">Blog Full Width</a>
                        </li>
                        <li>
                          <a href="blog-list-sidebar-left.html"
                            >Blog List Sidebar Left</a
                          >
                        </li>
                        <li>
                          <a href="blog-list-sidebar-right.html"
                            >Blog List Sidebar Right</a
                          >
                        </li>
                        <li>
                          <a href="blog-single-sidebar-left.html"
                            >Blog Single Sidebar left</a
                          >
                        </li>
                        <li>
                          <a href="blog-single-sidebar-right.html"
                            >Blog Single Sidebar Right</a
                          >
                        </li>
                      </ul>
                    </li>

                    <li v-if="false" class="has-dropdown">
                      <a href="#">Pages <i class="fa fa-angle-down"></i></a>
                      <!-- Sub Menu -->
                      <ul class="sub-menu">
                        <li><a href="faq.html">Frequently Questions</a></li>
                        <li>
                          <a href="privacy-policy.html">Privacy Policy</a>
                        </li>
                        <li><a href="404.html">404 Page</a></li>
                      </ul>
                    </li>

                    <li :class="{ on: $route.path === '/about' }">
                      <a href="/about" @click.prevent="goLink('/about')">{{
                        $t("bbs.about")
                      }}</a>
                    </li>

                    <li :class="{ on: $route.path === '/solutions' }">
                      <a href="/solutions" @click.prevent="goLink('/solutions')"
                        >{{ $t('bbs.solutions') }}</a>
                    </li>

                    <li :class="{ on: $route.path === '/meta' }">
                      <a href="/meta" @click.prevent="goLink('/meta')"
                        >{{ $t('admin.meta-world-explore') }}</a
                      >
                    </li>

                    <li :class="{ on: $route.path === '/marke_place' }">
                      <a
                        href="/marke_place"
                        @click.prevent="goLink('/marke_place')"
                        >{{ $t("bbs.market") }}</a
                      >
                    </li>
                    <li :class="{ on: $route.path === '/contact' }">
                      <a href="/contact" @click.prevent="goLink('/contact')"
                        >{{ $t('bbs.contactUs') }}</a
                      >
                      <!-- 切换时就是去 en-us tc-cn zh-cn 中找到对应key的value值进行渲染 懂了把 -->
                    </li>
                    <li :class="{ on: $route.path === '/forum' }">
                      <forum/>
                    </li>
                  </ul>
                </nav>
              </div>
              <!-- End 头部菜单 -->

              <!-- Start 头部 右侧小图标 -->
              <ul
                class="header-action-link action-color--black action-hover-color--golden"
              >
                <!-- Start 汉英切换 -->
                <li>
                  <el-tooltip
                    class="item"
                    effect="dark"
                    :content="$t('bbs.language')"
                    placement="bottom"
                  >
                    <a class="offcanvas-toggle" @click="setLanguage">
                      <i
                        :class="{
                          iconfont: true,
                          'icon-a-yuyanyingwen': language === 'en',
                          'icon-a-yuyanzhongwen': language === 'zh',
                          'icon-a-yuyanfanti': language === 'tc',
                        }"
                        style="font-size: 25px"
                      ></i>
                    </a>
                  </el-tooltip>
                </li>
                <!-- End  在这个文件  汉英切换 -->

                <!-- Start 收藏区 -->
                <li>
                  <el-tooltip
                    class="item"
                    effect="dark"
                    :content="$t('bbs.collect')"
                    placement="bottom"
                  >
                    <a class="offcanvas-toggle" @click="tip">
                      <i class="icon-heart"></i>
                    </a>
                  </el-tooltip>
                </li>
                <!-- End  收藏区 -->

                <li>
                  <el-tooltip
                    class="item"
                    effect="dark"
                    content="Cart"
                    placement="bottom"
                  >
                    <a class="offcanvas-toggle" @click="openCart">
                      <i class="icon-bag"></i>
                      <span v-if="cartList.length > 0" class="item-count">{{
                        cartList.length
                      }}</span>
                    </a>
                  </el-tooltip>
                </li>

                <!-- start 頭像區 -->
                <li class="avatar-li">
                  <el-popover
                    v-if="name"
                    placement="bottom"
                    trigger="hover"
                    width="200"
                  >
                    <el-menu class="nav-user">
                      <div class="user-name-box">
                        <i class="el-icon-s-custom"></i>
                        <span>{{ name }}</span>
                      </div>
                      <hr />
                      <el-menu-item
                        index="1"
                        class="nav-menu"
                        @click="goLink('/profile')"
                      >
                        <i class="el-icon-postcard"></i>
                        <span slot="title">User Panel</span>
                      </el-menu-item>

                      <el-menu-item index="2" class="nav-menu" @click="logout">
                        <i class="el-icon-sunset"></i>
                        <span slot="title">Log Out</span>
                      </el-menu-item>
                    </el-menu>
                    <el-avatar
                      slot="reference"
                      :src="avatar"
                      style="cursor: pointer"
                      size="medium"
                    >
                      <img
                        src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                      />
                    </el-avatar>
                  </el-popover>

                  <el-button
                    v-else
                    type="primary"
                    class="nav-login"
                    size="mini"
                    @click="goLogin"
                  >
                    <svg-icon
                      class="login-icon"
                      icon-class="Login-icon"
                    ></svg-icon>
                  </el-button>
                </li>

                <li>
                  <el-dropdown
                    v-if="!ismenuBar"
                    trigger="click"
                    szie="mini"
                    @command="setCommand"
                  >
                    <span
                      class="el-dropdown-link"
                      style="color: #fff; margin: 0 5px; font-size: 21px"
                    >
                      <i class="el-icon-s-fold el-icon--right"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown" :tabindex="0">
                      <el-dropdown-item
                        :class="{ select: 'home' == command }"
                        command="home"
                        @click.native="goLink('/')"
                        >{{ $t("bbs.home") }}</el-dropdown-item
                      >
                      <el-dropdown-item
                        :class="{ select: '1' == command }"
                        command="1"
                        @click.native="goLink('/about')"
                        >{{ $t("bbs.about") }}</el-dropdown-item
                      >
                      <el-dropdown-item
                        :class="{ select: '2' == command }"
                        command="2"
                        @click.native="goLink('/solutions')"
                        >{{ $t("bbs.solutions") }}</el-dropdown-item
                      >
                      <el-dropdown-item
                        :class="{ select: '3' == command }"
                        command="3"
                        @click.native="goLink('/meta')"
                        >{{ $t("admin.meta-world-explore") }}</el-dropdown-item
                      >
                      <el-dropdown-item
                        :class="{ select: '4' == command }"
                        command="4"
                        @click.native="goLink('/marke_place')"
                        >{{ $t("bbs.market") }}</el-dropdown-item
                      >
                      <el-dropdown-item
                        :class="{ select: '5' == command }"
                        command="5"
                        @click.native="goLink('/contact')"
                        >{{ $t("bbs.contact") }}</el-dropdown-item
                      >
                    </el-dropdown-menu>
                  </el-dropdown>
                </li>
                <!-- end 頭像區 -->
              </ul>
              <!-- End 头部 右侧小图标 -->
            </div>
          </div>
        </div>
      </div>
    </div>
    <CartDialog ref="CartDialog" />
  </header>
</template>

<script>
import forum from "@/views/bbs/forum.vue";
import { mapGetters } from "vuex";
import CartDialog from "@/components/CartDialog/CartDialog.vue";
export default {
  components: {
    CartDialog,
    forum
  },
  data() {
    return {
      isShowOrder: false, // 点击checkout弹窗
      ismenuBar: false, // 菜单
      command: "0",
    };
  },
  computed: {
    ...mapGetters(["name", "token", "cartList", "language", "avatar"]),
    isHide() {
      return this.$route.path === "/login" || this.$route.path === "/register";
    },
    subtotal: {
      get() {
        let price = 0.0;
        this.cartList.forEach((item) => {
          price += parseFloat(item.project_price);
        });
        return price.toFixed(2);
      },
    },
  },
  mounted() {
    this.handleResize();
    window.addEventListener("resize", this.handleResize);
  },
  methods: {
    setLanguage() {
      let language = "zh";
      switch (this.language) {
        case "zh":
          language = "tc";
          break;
        case "tc":
          language = "en";
          break;
        case "en":
          language = "zh";
          break;
        default:
          language = "zh";
          break;
      }
      this.$store.dispatch("locale/setLanguage", language);
      this.$i18n.locale = language;
    },
    handleResize() {
      this.ismenuBar = document.body.clientWidth > 1160;
    },
    // 点击上面购物袋
    openCart() {
      // 已登陆
      if (this.token) {
        this.$refs.CartDialog.open();
      } else this.$router.push("/login");
    },
    goLogin() {
      setTimeout(() => {
        this.$refs.CartDialog.close();
        this.$router.push("/login");
      }, 500);
    },
    goLink(path) {
      this.$router.push(path);
    },
    async logout() {
      await this.$store.dispatch("user/logout");
      this.$router.push(`/login?redirect=${this.$route.fullPath}`);
    },
    /** 下拉菜单设置高亮 */
    setCommand(command) {
      console.log(command);
      this.command = command;
    },
    tip() {
      this.$router.push("/wishlist");
    },
  },
};
</script>
<style src="@/assets/css/style.css" scoped></style>
<style src="@/assets/css/plugins/plugins.min.css" scoped></style>
<style src="@/assets/css/vendor/vendor.min.css" scoped></style>

<style lang="scss" scoped>
@media screen and (min-width: 1024px) {
  .logo {
    img {
      margin: 0 !important;
    }
    .svg-background {
      width: 15vw;
      height: 50px;
      background-image: url("../assets/images/pc/logo.svg");
      background-size: cover;
      background-repeat: no-repeat;
      background-position: center;
      // margin-top: -10vh;
    }
  }
}

@media screen and (max-width: 767px) {
  .logo {
    .svg-background {
      width: 40vw;
      height: 50px;
      background-image: url("../assets/images/pc/logo.svg");
      background-size: cover;
      background-repeat: no-repeat;
      background-position: center;
      // margin-top: -10px;
    }
  }
  .header-action-link {
    .avatar-li {
      margin-right: 0px !important;
    }
  }
}

@media screen and (min-width: 768px) and (max-width: 1023px) {
  .logo {
    img {
      height: 0.5rem;
      margin: 0;
    }
    .svg-background {
      width: 30vw;
      height: 80px;
      background-image: url("../assets/images/pc/logo.svg");
      background-size: cover;
      background-repeat: no-repeat;
      background-position: center;
      // margin-top: -10vh;
    }
  }
   .header-action-link {
    .avatar-li {
      margin-right: 15px !important;
    }
  }
}

/deep/ .select {
  background-color: #f57138;
  color: #fff !important;
}

/deep/ .el-avatar {
  img {
    width: 100%;
  }
}
/deep/ .el-dropdown-menu__item {
  color: #f57138;
  margin: 5px;
  border-radius: 5px;
  &:hover {
    background-color: #f57138;
    color: #fff;
  }
}

.cart-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.cart-empty-text {
  margin-top: 10px;
}
.cart-empty-text p {
  text-align: center;
}
.cart-empty-text p > a {
  color: #3ca3f4 !important;
}
.cart-empty-text p > a:hover {
  font-size: 16px;
  opacity: 0.7;
}
.cart-item {
  display: flex;
  border-bottom: solid 1px #eceff1;
}
.cart-item .cart-item-left {
  box-sizing: inherit;
  min-width: 250px;
  flex: 5;
  display: flex;
  margin: 16px 0;
}
.cart-item .cart-item-right > p {
  margin-top: 16px;
  color: #212121;
}
.cart-item-img {
  padding-right: 16px;
  min-width: 97px;
  box-sizing: inherit;
}
.cart-item-text {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 80px;
  width: 250px;
  box-sizing: inherit;
}
.cart-item-text .cart-item-text-2 {
  display: flex;
}
.cart-item-text .cart-item-text-2 > a {
  color: #9e9e9e;
  margin-right: 10px;
}
.cart-item-text .cart-item-text-2 > a:hover {
  color: #3ca3f4;
}
/deep/ .el-dialog__body {
  padding-top: 0px;
  overflow: auto;
}
/deep/ .el-dialog {
  position: absolute;
  left: 60%;
  width: 480px;
  height: 570px;
  border-radius: 4px;
}
/deep/.el-dialog__header {
  border-bottom: 1px solid rgba(0, 0, 0, 0.1) !important;
  height: 60px;
}
/deep/.el-dialog__header .el-dialog__title {
  font-size: 24px;
  font-weight: 700;
}
/deep/ .el-dialog__body {
  height: 366px;
}
/deep/ .el-dialog__footer {
  border-top: 1px solid rgba(0, 0, 0, 0.1) !important;
  height: 150px;
  padding: 24px;
}
.subtotal {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.subtotal .sub-text {
  font-size: 16px;
  font-weight: 700;
  text-align: left;
}
.subtotal .sub-price {
  font-size: 16px;
  font-weight: 700;
  float: right;
}
.tip {
  margin-top: 6px;
  text-align: left;
  color: #757575;
  font-size: 14px;
  font-weight: 400;
}
.cart-button {
  margin-top: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
}
.cart-button > .btn-1,
.cart-button > .btn-2 {
  height: 40px;
  width: 45%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: 700;
}
.cart-button > .btn-1 {
  border: 1px solid #3ca3f4;
  color: #3ca3f4;
  border-radius: 4px;
  margin-right: 15px;
}

.cart-button > .btn-1,
.cart-button > .btn-2:hover {
  opacity: 0.8;
}

.cart-button > .btn-2 {
  background-color: #3ca3f4;
  border: 1px solid #3ca3f4;
  border-radius: 4px;
  color: #fff;
}

.header_btn {
  cursor: pointer;
  &:hover {
    color: #000;
  }
}

.header-section {
  width: 100%;
  position: absolute;
  top: 0;
  // border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
  z-index: 999;
  // background-color: red;
}

.menu-color--black nav > ul > li > a:hover {
  color: #fff !important;
  background-color: #be571789 !important;
}
.on > a {
  background-color: #be5717b9 !important;
}
.header-action-link .offcanvas-toggle {
  color: #fff !important;
}
.nav-user {
  border: none;
}
.nav-user > li > span,
.nav-user > li > i {
  color: #f57138;
}
.nav-user > li > i {
  font-weight: bold;
}
.nav-menu {
  border-radius: 10px;
}
.nav-menu:hover {
  background-color: #f57138;
}
.nav-menu:hover * {
  color: #fff !important;
}
.login-icon {
  width: 25px !important;
  height: 23px !important;
}
.nav-login {
  // width: 60px !important;
  // height: 50px !important;
  // left: 20px;
  // &:hover{
  //   background-color:black;
  // }
  padding: 0 !important;
  border-radius: 20px !important;
  background-color: transparent !important;
  border-color: #f57138;
  color: #f57138;
  // box-shadow: 0 5px 2px #243c479f;
  // transition: all 0.5s;
  // box-sizing: border-box;
  // transform: scale(0.8);
}
// .nav-login:active {
//   box-shadow: 0 4px 0 #243c47;
//   transform: translateY(4px);
// }

.user-name-box {
  width: 100%;
  min-height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background-image: linear-gradient(to right top, #ffca65, #ff7344);
  cursor: default;
}
.user-name-box i,
.user-name-box span {
  color: #fff;
  font-size: 14px;
  margin-right: 10px;
}
.item-count {
  background-color: #f57138;
  color: #fff;
}

.order-dialog {
  /deep/ .el-dialog {
    position: relative;
    left: 0;
  }

  /deep/ .el-dialog__footer {
    height: auto;
    padding: 0;
  }

  /deep/.el-dialog__header {
    border-bottom: unset !important;
  }

  /deep/.el-dialog__body {
    height: unset;
  }
  .order-content {
    overflow: auto;
    max-height: 400px;
    margin-top: 20px;
    .order-header {
      padding-bottom: 20px;
      display: flex;
      border-bottom: 1px solid rgba(0, 0, 0, 0.1);
      justify-content: space-between;
      .order-header-left {
        p {
          margin-bottom: 5px;
          font-size: 10px;
        }
      }
      .order-header-right {
        p {
          margin-bottom: 5px;
        }
        width: 200px;
      }
    }
    .order-pay {
      margin-top: 20px;
      display: flex;
      margin-bottom: 20px;
      align-items: center;
      p {
        margin-right: 5px;
        margin-bottom: 0px;
      }
    }
    .order-table {
      /deep/ .el-table {
        font-size: 12px;
      }
    }
    .order-total {
      margin-top: 15px;
      margin-right: 60px;
      float: right;
      .order-total-content {
        margin-bottom: 5px;
        p {
          font-size: 12px;
          margin-bottom: 0px;
        }
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 170px;
      }
    }
  }
}
.order_content {
  .order_content_top {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-top: 20px;
  }
  .order_content_bottom {
    &-content {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      margin-top: 20px;
      &-left {
        display: flex;
        align-items: center;
        span {
          margin-left: 10px;
        }
      }
    }
  }
}
</style>
