import axios from "axios";

const state = () => ({
  products: [],
  isLoading: false,
  addToCartResult: "",
  isShowCartDropdown: false,
});

const getters = {
  totalItems(state) {
    if(state.products === null || state.products === undefined ) return 0
    return state.products.reduce(
      (total, product) => total + product.quantity,
      0
    );
  },

  subTotal(state) {
    if(state.products === null || state.products === undefined ) return 0
    return state.products.reduce(
      (totalPrice, product) => totalPrice + product.quantity * product.product.price,
      0
    );
  },
};

const actions = {
  async updateCart({commit}, {userId, product, quantity, replace = "", isPaid = 0}) {
      try {
        const res = await axios.post("order/upsert", {user_id: userId, product_id: product.id, quantity: quantity, replace: replace, is_paid: isPaid})
        console.log("ok", res)
        if (replace === "") {
          commit("addProductToCart", {product: product, quantity: quantity})
        }
      } catch (e) {
        console.log(e)
      }
  },
  async getCartById({commit}, id) {
    try {
      const res = await axios.get(`/orders/user/${id}`)
      commit("CALL_CART_FROM_SV", res.data)
      console.log(res)
    } catch (e) {
      console.log(e)
    }
  },
  async deleteOrderItem({commit}, id) {
    try {
      const res = axios.delete(`/orders/${id}`)
      console.log(res)
      commit("REMOVE_ORDER", id)
    } catch (e) {
      console.log(e)
    }
  }

};

const mutations = {
  CALL_CART_FROM_SV(state, cart) {
      state.products = cart
  },
  SET_PRODUCTS(state, products) {
    state.products = products;
  },
  setShowCartDropdown(state, status) {
    state.isShowCartDropdown = status;
  },

  updateProductQuantity(state, { productId, value }) {
    const product = state.products.find((p) => p.id === productId);
    const index = state.products.findIndex(item => item.id === productId)
    value = Number(value);

    if (value > 1) {
      product.quantity = value;
    } else {
      state.products.splice(index,1)
    }

    product.totalPrice = product.price * product.quantity;
  },
  addProductToCart(state, {product, quantity}) {
    console.log(quantity)
    let isExist = state.products?.findIndex(item => item.product.id === product.id)
    if (isExist === undefined) {
      console.log("isExist === undefined")
      isExist = -1
      state.products = []
    }
    if(quantity === 0 ) return
    if (isExist === -1) {
      state.products.push({product: product, quantity: quantity});
    } else {
      console.log(isExist)
      isExist = state.products.findIndex(item => item.product.id === product.id)
      state.products[isExist].quantity = state.products[isExist].quantity + quantity
    }
  },
  REMOVE_ORDER(state, id) {
    state.products = state.products.filter(product => product.id !== id)
  },
  REMOVE_CART(state) {
    state.products = []
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
