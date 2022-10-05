package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReverseString", func() {
	When("input str contains 'Hello World'", func() {
		It("should return 'd_l_r_o_W o_l_l_e_H'", func() {
			Expect(main.ReverseString("Hello World")).To(Equal("d_l_r_o_W o_l_l_e_H"))
		})
	})

	When("input str contains 'I am a student'", func() {
		It("should return 't_n_e_d_u_t_s a m_a I'", func() {
			Expect(main.ReverseString("I am a student")).To(Equal("t_n_e_d_u_t_s a m_a I"))
		})
	})

	When("input str contains 'I am a stranger'", func() {
		It("should return 'r_e_g_n_a_r_t_s a m_a I'", func() {
			Expect(main.ReverseString("I am a stranger")).To(Equal("r_e_g_n_a_r_t_s a m_a I"))
		})
	})

	When("input str contains 'I am a student and Stranger And Person to Improvement'", func() {
		It("should return 't_n_e_m_e_v_o_r_p_m_I o_t n_o_s_r_e_P d_n_A r_e_g_n_a_r_t_S d_n_a t_n_e_d_u_t_s a m_a I'", func() {
			Expect(main.ReverseString("I am a student and Stranger And Person to Improvement")).To(Equal("t_n_e_m_e_v_o_r_p_m_I o_t n_o_s_r_e_P d_n_A r_e_g_n_a_r_t_S d_n_a t_n_e_d_u_t_s a m_a I"))
		})
	})

	When("input str is letter of lorem ipsum", func() {
		It("should return reverse letter of lorem ipsum", func() {
			var lorem = "Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Volutpat lacus laoreet non curabitur Amet risus nullam eget felis eget nunc Amet risus nullam eget felis eget Consectetur purus ut faucibus pulvinar elementum Suspendisse ultrices gravida dictum fusce ut placerat orci nulla pellentesque Eu nisl nunc mi ipsum Ac turpis egestas maecenas pharetra convallis Egestas tellus rutrum tellus pellentesque eu tincidunt tortor Nunc sed velit dignissim sodales ut Ut venenatis tellus in metus vulputate eu scelerisque felis Feugiat pretium nibh ipsum consequat nisl vel Porta non pulvinar neque laoreet suspendisse interdum consectetur Sed elementum tempus egestas sed sed risus Blandit volutpat maecenas volutpat blandit aliquam etiam erat Porta non pulvinar neque laoreet suspendisse interdum Vel fringilla est ullamcorper eget nulla facilisi etiam dignissim diam Nisl suscipit adipiscing bibendum est ultricies integer quis auctor Vitae tempus quam pellentesque nec nam aliquam sem Malesuada bibendum arcu vitae elementum curabitur vitae nunc sed Quam elementum pulvinar etiam non quam lacus suspendisse Nunc mattis enim ut tellus elementum sagittis vitae et Ultrices eros in cursus turpis massa tincidunt Morbi tincidunt ornare massa eget egestas purus viverra Massa ultricies mi quis hendrerit dolor magna eget Ut eu sem integer vitae justo eget magna fermentum iaculis Diam sit amet nisl suscipit adipiscing Convallis aenean et tortor at risus viverra Nulla posuere sollicitudin aliquam ultrices sagittis orci a Quam viverra orci sagittis eu volutpat odio facilisis mauris Libero volutpat sed cras ornare arcu dui vivamus Hendrerit dolor magna eget est lorem ipsum dolor sit amet Congue eu consequat ac felis donec et odio Etiam tempor orci eu lobortis elementum nibh tellus At auctor urna nunc id cursus Velit egestas dui id ornare arcu odio ut Blandit massa enim nec dui nunc mattis enim ut tellus Risus quis varius quam quisque Interdum velit laoreet id donec ultrices tincidunt arcu non sodales Vel risus commodo viverra maecenas"
			var expected = "s_a_n_e_c_e_a_m a_r_r_e_v_i_v o_d_o_m_m_o_c s_u_s_i_r l_e_V s_e_l_a_d_o_s n_o_n u_c_r_a t_n_u_d_i_c_n_i_t s_e_c_i_r_t_l_u c_e_n_o_d d_i t_e_e_r_o_a_l t_i_l_e_v m_u_d_r_e_t_n_I e_u_q_s_i_u_q m_a_u_q s_u_i_r_a_v s_i_u_q s_u_s_i_R s_u_l_l_e_t t_u m_i_n_e s_i_t_t_a_m c_n_u_n i_u_d c_e_n m_i_n_e a_s_s_a_m t_i_d_n_a_l_B t_u o_i_d_o u_c_r_a e_r_a_n_r_o d_i i_u_d s_a_t_s_e_g_e t_i_l_e_V s_u_s_r_u_c d_i c_n_u_n a_n_r_u r_o_t_c_u_a t_A s_u_l_l_e_t h_b_i_n m_u_t_n_e_m_e_l_e s_i_t_r_o_b_o_l u_e i_c_r_o r_o_p_m_e_t m_a_i_t_E o_i_d_o t_e c_e_n_o_d s_i_l_e_f c_a t_a_u_q_e_s_n_o_c u_e e_u_g_n_o_C t_e_m_a t_i_s r_o_l_o_d m_u_s_p_i m_e_r_o_l t_s_e t_e_g_e a_n_g_a_m r_o_l_o_d t_i_r_e_r_d_n_e_H s_u_m_a_v_i_v i_u_d u_c_r_a e_r_a_n_r_o s_a_r_c d_e_s t_a_p_t_u_l_o_v o_r_e_b_i_L s_i_r_u_a_m s_i_s_i_l_i_c_a_f o_i_d_o t_a_p_t_u_l_o_v u_e s_i_t_t_i_g_a_s i_c_r_o a_r_r_e_v_i_v m_a_u_Q a i_c_r_o s_i_t_t_i_g_a_s s_e_c_i_r_t_l_u m_a_u_q_i_l_a n_i_d_u_t_i_c_i_l_l_o_s e_r_e_u_s_o_p a_l_l_u_N a_r_r_e_v_i_v s_u_s_i_r t_a r_o_t_r_o_t t_e n_a_e_n_e_a s_i_l_l_a_v_n_o_C g_n_i_c_s_i_p_i_d_a t_i_p_i_c_s_u_s l_s_i_n t_e_m_a t_i_s m_a_i_D s_i_l_u_c_a_i m_u_t_n_e_m_r_e_f a_n_g_a_m t_e_g_e o_t_s_u_j e_a_t_i_v r_e_g_e_t_n_i m_e_s u_e t_U t_e_g_e a_n_g_a_m r_o_l_o_d t_i_r_e_r_d_n_e_h s_i_u_q i_m s_e_i_c_i_r_t_l_u a_s_s_a_M a_r_r_e_v_i_v s_u_r_u_p s_a_t_s_e_g_e t_e_g_e a_s_s_a_m e_r_a_n_r_o t_n_u_d_i_c_n_i_t i_b_r_o_M t_n_u_d_i_c_n_i_t a_s_s_a_m s_i_p_r_u_t s_u_s_r_u_c n_i s_o_r_e s_e_c_i_r_t_l_U t_e e_a_t_i_v s_i_t_t_i_g_a_s m_u_t_n_e_m_e_l_e s_u_l_l_e_t t_u m_i_n_e s_i_t_t_a_m c_n_u_N e_s_s_i_d_n_e_p_s_u_s s_u_c_a_l m_a_u_q n_o_n m_a_i_t_e r_a_n_i_v_l_u_p m_u_t_n_e_m_e_l_e m_a_u_Q d_e_s c_n_u_n e_a_t_i_v r_u_t_i_b_a_r_u_c m_u_t_n_e_m_e_l_e e_a_t_i_v u_c_r_a m_u_d_n_e_b_i_b a_d_a_u_s_e_l_a_M m_e_s m_a_u_q_i_l_a m_a_n c_e_n e_u_q_s_e_t_n_e_l_l_e_p m_a_u_q s_u_p_m_e_t e_a_t_i_V r_o_t_c_u_a s_i_u_q r_e_g_e_t_n_i s_e_i_c_i_r_t_l_u t_s_e m_u_d_n_e_b_i_b g_n_i_c_s_i_p_i_d_a t_i_p_i_c_s_u_s l_s_i_N m_a_i_d m_i_s_s_i_n_g_i_d m_a_i_t_e i_s_i_l_i_c_a_f a_l_l_u_n t_e_g_e r_e_p_r_o_c_m_a_l_l_u t_s_e a_l_l_i_g_n_i_r_f l_e_V m_u_d_r_e_t_n_i e_s_s_i_d_n_e_p_s_u_s t_e_e_r_o_a_l e_u_q_e_n r_a_n_i_v_l_u_p n_o_n a_t_r_o_P t_a_r_e m_a_i_t_e m_a_u_q_i_l_a t_i_d_n_a_l_b t_a_p_t_u_l_o_v s_a_n_e_c_e_a_m t_a_p_t_u_l_o_v t_i_d_n_a_l_B s_u_s_i_r d_e_s d_e_s s_a_t_s_e_g_e s_u_p_m_e_t m_u_t_n_e_m_e_l_e d_e_S r_u_t_e_t_c_e_s_n_o_c m_u_d_r_e_t_n_i e_s_s_i_d_n_e_p_s_u_s t_e_e_r_o_a_l e_u_q_e_n r_a_n_i_v_l_u_p n_o_n a_t_r_o_P l_e_v l_s_i_n t_a_u_q_e_s_n_o_c m_u_s_p_i h_b_i_n m_u_i_t_e_r_p t_a_i_g_u_e_F s_i_l_e_f e_u_q_s_i_r_e_l_e_c_s u_e e_t_a_t_u_p_l_u_v s_u_t_e_m n_i s_u_l_l_e_t s_i_t_a_n_e_n_e_v t_U t_u s_e_l_a_d_o_s m_i_s_s_i_n_g_i_d t_i_l_e_v d_e_s c_n_u_N r_o_t_r_o_t t_n_u_d_i_c_n_i_t u_e e_u_q_s_e_t_n_e_l_l_e_p s_u_l_l_e_t m_u_r_t_u_r s_u_l_l_e_t s_a_t_s_e_g_E s_i_l_l_a_v_n_o_c a_r_t_e_r_a_h_p s_a_n_e_c_e_a_m s_a_t_s_e_g_e s_i_p_r_u_t c_A m_u_s_p_i i_m c_n_u_n l_s_i_n u_E e_u_q_s_e_t_n_e_l_l_e_p a_l_l_u_n i_c_r_o t_a_r_e_c_a_l_p t_u e_c_s_u_f m_u_t_c_i_d a_d_i_v_a_r_g s_e_c_i_r_t_l_u e_s_s_i_d_n_e_p_s_u_S m_u_t_n_e_m_e_l_e r_a_n_i_v_l_u_p s_u_b_i_c_u_a_f t_u s_u_r_u_p r_u_t_e_t_c_e_s_n_o_C t_e_g_e s_i_l_e_f t_e_g_e m_a_l_l_u_n s_u_s_i_r t_e_m_A c_n_u_n t_e_g_e s_i_l_e_f t_e_g_e m_a_l_l_u_n s_u_s_i_r t_e_m_A r_u_t_i_b_a_r_u_c n_o_n t_e_e_r_o_a_l s_u_c_a_l t_a_p_t_u_l_o_V a_u_q_i_l_a a_n_g_a_m e_r_o_l_o_d t_e e_r_o_b_a_l t_u t_n_u_d_i_d_i_c_n_i r_o_p_m_e_t d_o_m_s_u_i_e o_d d_e_s t_i_l_e g_n_i_c_s_i_p_i_d_a r_u_t_e_t_c_e_s_n_o_c t_e_m_a t_i_s r_o_l_o_d m_u_s_p_i m_e_r_o_L"

			Expect(main.ReverseString(lorem)).To(Equal(expected))
		})
	})
})
