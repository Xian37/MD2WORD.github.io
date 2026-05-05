# homework_cleaned

**日期：2026-04-27**

---

### I. 緒論 (Introduction)
本章節探討氮化鎵 (GaN) 在高頻電源轉換器中的應用優勢及其物理限制帶來的設計挑戰。

* **GaN 物理特性與應用優勢：** GaN 是一種高電子移動率電晶體 (High Electron Mobility Transistor, HEMT)，具備低導通電阻 (Low on-resistance)、低寄生電容 (Low parasitic capacitances) 與高崩潰電壓 (High breakdown voltage) 等特性 。相較於傳統矽 (Si) 元件，GaN 能在高切換頻率下提供更高效的電源轉換，非常適合對空間與快速響應有嚴格要求的第五代行動通訊 (5G) 電信電源與工業電源 。
* **設計挑戰一：逆向傳導損耗 (Reverse Conduction Loss)：** 高至低停死時間 (High-to-low dead-time) 定義為切換節點電壓 $V_x$ 下降邊緣與低側開關閘源極電壓 $V_{gs,QL}$ 上升邊緣之間的時段 。由於 $V_x$ 在高至低停死時間的電壓斜率 (Slew rate) 由電感電流決定，當負載增加時，停死時間會隨之拉長 。GaN 元件缺乏本體二極體 (Body diode)，其逆向偏壓電壓降 (大於 1.5V) 遠高於矽元件 (典型值 0.7V)，且該電壓降會隨負載增加而上升 。這導致在重載下產生極大的逆向傳導損耗，因此必須將停死時間極小化 。
* **設計挑戰二：閘極驅動電壓下降 (Gate Voltage Drop)：** GaN 僅具有 n 型結構，因此高側開關必須依賴拔靴電路 (Bootstrap circuit) 提供合適的閘極超驅動電壓 (Gate overdrive voltage) 。GaN 元件的導通電阻與閘源極電壓 $V_{gs}$ 成反比 。若使用晶片外 (Off-chip) 大電容雖能減少電壓降，但打線 (Bonding wire) 與印刷電路板 (PCB) 造成的額外寄生電感，會導致 $V_{gs}$ 產生電壓突波 (Voltage spike)，可能引發增強型 (E-mode) GaN 元件的閘極崩潰 (Gate breakdown) 。因此，必須使用晶片內建電容 (On-chip capacitor)，但其小面積會帶來嚴重的電壓下降問題 。

---

### II. 系統架構概述 (System Overview)
提出的整合型驅動器 IC 架構包含兩大核心區塊，用於驅動兩個 E-mode EPC8009 GaN FETs ($Q_H$ 與 $Q_L$) 。

* **Bang-Bang 停死時間控制 (Bang-Bang Dead-Time Control)：** 包含比較器感測 (Comparator sensing)、延遲時間調變器 (Delay time modulator) 與停死時間產生器 (Dead-time generator)，負責藉由鎖定切換節點 $V_x$ 與低側閘極 $V_{gs,QL}$ 波形，將高至低停死時間最小化以減少逆向傳導損耗 。
* **電荷共享拔靴電路 (Charge Sharing Bootstrap Circuit)：** 採用一顆低壓與一顆高壓晶片內建拔靴電容，透過開關進行電荷共享，在減少晶片面積的同時，為高側驅動器提供足夠高的驅動電壓 $V_{boot}$，並避免寄生電感導致的 $V_{gs}$ 切換突波 。

---

### III. Bang-Bang 停死時間控制 (Bang-Bang Dead-Time Control)

#### A. 運作原理 (Operation Principle)
由於低至高停死時間 (Low-to-high dead-time) 不會隨負載改變，為降低控制複雜度，本設計僅針對高至低停死時間進行動態調整 。
* **訊號感測與邏輯判斷：** 理想的控制是使 $V_x$ 下降邊緣與 $V_{gs,QL}$ 導通時機同步 。由於 GaN 元件在達到臨界電壓 $V_{th}$ 時導通，系統將 $V_{gs,QL}$ 與 $V_{th}$ 饋入比較器以獲得真實的導通時機，同時 $V_x$ 也饋入另一比較器 。比較器輸出 $V_{X\_OUT}$ 與 $LG\_OUT$ 送至停死時間偵測區塊 。
* **UP 與 DN 訊號產生：** * **停死時間 (Dead-time)：** 若 $V_x$ 下降邊緣早於 $V_{gs,QL}$ 導通，表示有停死時間，此時當 $V_{X\_OUT}$、$LG\_OUT$ 與反相工作週期 $Duty\_inv$ 皆為高電位時，輸出 UP 訊號 。
    * **重疊 (Overlap)：** 若 $V_x$ 下降邊緣晚於 $V_{gs,QL}$ 導通，表示發生重疊，此時當 $V_{X\_OUT}$ 與 $LG\_OUT$ 皆為低電位且 $Duty\_inv$ 為高電位時，輸出 DN 訊號 。
* **Bang-Bang 控制與延遲產生：** 由於 UP 與 DN 脈波過窄會導致控制電壓 $V_c$ 補償緩慢，系統引入 Bang-Bang 控制，將訊號延伸為 $UP\_OUT$ 與 $DN\_OUT$，送至電荷泵 (Charge pump) 快速調整 $V_c$ 。
* **電壓控制延遲線 (Voltage-controlled delay line)：** $V_c$ 藉由控制電晶體 $M_{P3}$ 的導通電阻來改變延遲時間 。當 $V_c$ 為高時，$M_{P3}$ 電阻大，RC 時間常數近似於 $M_{P1}$ 導通電阻與 $V_A$ 節點寄生電容的乘積 。當 $V_c$ 為低時，$M_{P3}$ 電阻小，RC 時間常數需加入 $C_d$，導致延遲時間增加 。

#### B. 訊號感測技術細節 (Signal Sensing)
高頻 GaN 轉換器切換時會產生劇烈的電壓突波與震盪 (Ringing)，容易導致停死時間誤判 。
* **SR 閂鎖器 (SR Latch)：** 在 $V_x$ 比較器輸出端 $V_{COM}$ 後方加入 SR 閂鎖器，成功在感測到 $V_x$ 下降邊緣後將訊號箝位 (Clamp)，避免後續震盪干擾 。
* **克耳文感測 (Kelvin Sensing)：** 大且高轉換率 (High-slew-rate) 的驅動電流 $I_g$ 流經 IC 打線寄生電感時，會在低側閘極 (LG) 訊號上產生巨大突波 。系統採用獨立的克耳文感測路徑測量 $LG\_sense$，由於感測路徑上的 $I_{sense}$ 具備極低的 di/dt，因此能獲得乾淨無突波的波形以供精確比對 。

---

### IV. 電荷共享拔靴電路 (Charge Sharing Bootstrap Circuit)

#### A. 核心概念與公式推導 (Concept & Derivations)
為了在小晶片面積內減少拔靴電容電壓下降量，系統利用高壓 (HV) 電容 $C_{boot2}$ 儲存 12V 輸入電壓 $V_{IN}$ 的能量 。
定義等效閘極電荷 $Q_g$ 在開關導通時的位移量與電容的關係為：
$$Q_g = C_{boot} \cdot V_{dip} + C_{boot2} \cdot V_{dip2}$$ 

其中 $C_{boot}$ 為 5V 拔靴電容，$C_{boot2}$ 為 12V 拔靴電容；$V_{dip}$ 與 $V_{dip2}$ 分別為兩者的電壓下降量 。
由於在導通 (Duty ON) 期間，$C_{boot2}$ 會與 $C_{boot}$ 共享電荷，兩者的最終電壓將達到相等，因此 $V_{dip2}$ 可改寫為：
$$V_{dip2} = V_{Cboot2} - (V_{Cboot} - V_{dip})$$

將上式代入第一式，可得：
$$Q_g = C_{boot} \cdot V_{dip} + C_{boot2} \cdot [V_{Cboot2} - (V_{Cboot} - V_{dip})]$$

最終推導出 $C_{boot}$ 的電壓下降量方程式：
$$V_{dip} = \frac{Q_g - C_{boot2} \cdot (V_{Cboot2} - V_{Cboot})}{C_{boot} + C_{boot2}}$$

此公式表明，由於 $V_{Cboot2}$ (受 $V_{IN}$ 充電) 恆大於 $V_{Cboot}$ (受 $V_{DRV}$ 充電)，提出的拔靴電路能有效降低驅動電壓的下降量 。

#### B. 電路實作與保護機制 (Circuit Implementation and Operation)
* **充電階段 (Duty OFF)：** 此時低側開關導通，$V_x$ 接近接地 。$V_{DRV}$ 透過 $M_A$ 與 $D_A$ 對 $C_{boot}$ 充電；$V_{IN}$ 透過二極體 $D_B$ 對 $C_{boot2}$ 充電；$V_{DRV}$ 透過 $M_C$ 與 $D_C$ 對 $C_{boot3}$ 充電 。採用高壓 NMOS 實作 $M_{SW1}$ 與 $M_{SW2}$ 開關以承受超過 5V 的跨壓 。
* **共享階段 (Duty ON)：** 此時高側開關導通，$V_x$ 被拉升至 $V_{IN}$ 。$C_{boot2}$ 會透過高速路徑 $M_{SW2}$ 直接供電給 $V_{GH}$ 。當 $V_{GH}$ 升至 $V_{boot}$ 減去 $M_{SW2}$ 臨界電壓時，$M_{SW2}$ 會關閉以避免 $Q_H$ 閘極過度應力 (Overstress) 。同時，位準轉換器輸出 (LSF_OUT) 配合 $V_{Cboot3}$ 會開啟 $M_{SW1}$，使 $C_{boot2}$ 與 $C_{boot}$ 進行電荷共享 。
* **停死時間保護 (Dead-Time Protection)：** 在停死時間，$V_x$ 會降至約 -2V，這會導致 $C_{boot}$ 與 $C_{boot3}$ 過度充電至 7V，進而擊穿驅動 IC 內的 MOSFET 。為解決此問題，引入訊號 $P_{gate}$ 與 $P_{gate2}$ 在停死時間強迫關閉 $M_A$ 與 $M_C$ 。為了避免 $M_A$ 與 $M_C$ 汲極端產生浮接節點 (Floating node)，額外加入 $M_D$ 與 $M_E$ 以提供低電位 。

#### C. 設計考量 (Design Consideration)
由於真實電路中 $Q_g$ 存在製程變異，且切換節點 $V_x$ 在停死時間會隨負載變動，導致 $V_{Cboot2}$ 也隨之改變 。需計算極端情況：
最大電壓下降量：
$$V_{dip(max)} = \frac{Q_{g(max)} - C_{boot2}(V_{Cboot2(min)} - V_{Cboot})}{C_{boot} + C_{boot2}}$$

最小電壓下降量：
$$V_{dip(min)} = \frac{Q_{g(min)} - C_{boot2}(V_{Cboot2(max)} - V_{Cboot})}{C_{boot} + C_{boot2}}$$

依據上述公式與參數，最終選定 $C_{boot2}$ 為 60 pF，使 $C_{boot}$ 電壓下降量的平均值趨近於零 。

---

### V. 實驗驗證與數據分析 (Experimental Verification)

* **晶片與系統參數：** 採用 TSMC 0.25-µm HV BCD 製程製造，總面積為 1.5 $mm^2$ 。內建 0.2 nF 總電容，包含 100 pF $C_{boot}$、60 pF $C_{boot2}$ 以及 40 pF $C_{boot3}$ 。轉換器操作於 10 MHz 頻率、12V 輸入、5V 輸出，負載範圍為 0.2A 至 0.65A 。
* **停死時間鎖定表現：** 傳統固定停死時間控制在輕載時 $V_x$ 下降時間為 9 ns (停死時間 1.8 ns)，重載時下降時間為 2.6 ns (停死時間暴增至 9.2 ns) 。使用提出的 Bang-Bang 控制後，不論輕載或重載，皆成功將停死時間死鎖於最小 400 ps (0.4 ns) 至最大 1.4 ns 的區間 。
* **動態響應與穩定性：** 控制電壓 $V_c$ 的整定速度 (Settling speed) 達到每微秒 0.941V 。在負載瞬態變化期間 (1 µs 內由 0.65A 降至 0.2A)，由於 $V_x$ 下降斜率與電感電流相關，不會發生上下管直通 (Shoot-through) 或硬切換 (Hard-switching) 。
* **驅動電壓維持：** 在輕重載條件下，高側閘源極電壓皆穩定維持在 4.365V 。此設計僅使用了傳統設計 45% 的電容量，即將電壓下降量控制在 0.6V 以內 (5V 理論值減去 4.365V 約為 0.635V) 。
* **效率與損耗細分 (Loss Breakdown)：** * 由於電荷共享拔靴電路提升了驅動電壓，單獨使效率改善了 1.5%，並將重載時的 GaN HEMT 傳導損耗佔比由 19% 降至 10% 。
    * 由於停死時間的大幅縮短，重載時的停死時間損耗佔比由 22.31% 劇降至 5% 。
    * 在 0.65A 重載條件下，結合兩項技術後的整體效率達到了 89.9%，較傳統控制提升了 5% 。