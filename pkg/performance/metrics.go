package performance

import (
	"strconv"
	"testing"

	"github.com/maistra/maistra-test-tool/pkg/util"
)

func TestXDSPushes(t *testing.T) {
	util.Log.Info("** TEST: TestXDSPushes")
	xdsPushCount, err := getMetricPrometheusMesh("xds_ppctc")
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	xdsPushTime, err := getMetricPrometheusMesh("xds_ppctb")
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}

	xdsPushCountValue, err := parseResponse([]byte(xdsPushCount))
	xdsPushTimeValue, err := parseResponse([]byte(xdsPushTime))

	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}

	util.Log.Info(" If xdsPushCount and xdsPushTime are equal - OK")
	if xdsPushCountValue[0] != xdsPushTimeValue[0] {
		t.Errorf("xdsPushCount (%v) and xdsPushTime (%v) are not equal", xdsPushCountValue, xdsPushTimeValue)
	}
}

func TestIstiodMem(t *testing.T) {
	util.Log.Info("** TEST: TestIstiodMem")
	istiodMem, err := getMetricPrometheusOCP("istiod_mem", nil)
	istiodMemValue, err := parseResponse([]byte(istiodMem))
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}

	// Transform values to integers and compare them in bytes
	istiodMemValueInt, err := strconv.Atoi(istiodMemValue[0])
	istiodAcceptanceMemIntBytes, err := strconv.Atoi(istiodAcceptanceMem)

	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}

	istiodAcceptanceMemIntBytes = istiodAcceptanceMemIntBytes * bytesToMegaBytes

	util.Log.Info(" If istiodMem is lower than ", istiodAcceptanceMem, "MB")
	if istiodMemValueInt > istiodAcceptanceMemIntBytes {
		t.Errorf("Istiod Memory Value is %v. Want something lower than %v", istiodMemValueInt, istiodAcceptanceMemIntBytes)
	}
}

func TestIstiodCpu(t *testing.T) {
	util.Log.Info("** TEST: TestIstiodCpu")
	istiodCpu, err := getMetricPrometheusOCP("istiod_cpu", nil)
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	istiodCpuValue, err := parseResponse([]byte(istiodCpu))

	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}

	istiodAcceptanceCpuInt, err := strconv.ParseFloat(istiodAcceptanceCpu, 32)
	istiodCpuValueInt, err := strconv.ParseFloat(istiodCpuValue[0], 32)

	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}

	util.Log.Info(" If istiodCpu is lower than ", istiodAcceptanceCpu)
	if istiodCpuValueInt > istiodAcceptanceCpuInt {
		t.Errorf("Istiod CPU Value is %v. Want something lower than %v", istiodCpuValueInt, istiodAcceptanceCpuInt)
	}
}

func TestIstioProxiesMem(t *testing.T) {
	util.Log.Info("** TEST: TestIstioProxiesMem")

	meshPods, err := getMeshProxies()
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	for name, namespace := range meshPods {
		prometheusAPIMapParams["istio-proxy-pod-name"] = name
		prometheusAPIMapParams["istio-proxy-pod-ns"] = namespace
		istiodMem, err := getMetricPrometheusOCP("istio_proxies_mem_custom", prometheusAPIMapParams)
		if err != nil {
			util.Log.Error(err)
			t.Error(err)
			t.FailNow()
		}
		util.Log.Info(istiodMem)
		util.Log.Info(" If IstioProxiesMem is lower than ", istioProxiesAcceptanceMem)
	}
}

func TestIstioProxiesCpu(t *testing.T) {
	util.Log.Info("** TEST: TestIstioProxiesCpu")

	meshPods, err := getMeshProxies()
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	for name, namespace := range meshPods {
		prometheusAPIMapParams["istio-proxy-pod-name"] = name
		prometheusAPIMapParams["istio-proxy-pod-ns"] = namespace
		istiodCpu, err := getMetricPrometheusOCP("istio_proxies_cpu_custom", prometheusAPIMapParams)
		if err != nil {
			util.Log.Error(err)
			t.Error(err)
			t.FailNow()
		}
		util.Log.Info(istiodCpu)
		util.Log.Info(" If istioProxiesCpu is lower than ", istioProxiesAcceptanceCpu)
	}
}

func TestIstioIngressProxiesMem(t *testing.T) {
	util.Log.Info("** TEST: TestIstioIngressProxiesMem")

	meshPods, err := getMeshIngressProxies()
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	for name, namespace := range meshPods {
		prometheusAPIMapParams["istio-proxy-pod-name"] = name
		prometheusAPIMapParams["istio-proxy-pod-ns"] = namespace
		istioIngressProxyMem, err := getMetricPrometheusOCP("istio_proxies_mem_custom", prometheusAPIMapParams)
		if err != nil {
			util.Log.Error(err)
			t.Error(err)
			t.FailNow()
		}
		util.Log.Info(istioIngressProxyMem)
		util.Log.Info(" If istioIngressProxiesAcceptanceMem is lower than ", istioIngressProxiesAcceptanceMem)
	}
}

func TestIstioIngressProxiesCpu(t *testing.T) {
	util.Log.Info("** TEST: TestIstioIngressProxiesCpu")

	meshPods, err := getMeshIngressProxies()
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	for name, namespace := range meshPods {
		prometheusAPIMapParams["istio-proxy-pod-name"] = name
		prometheusAPIMapParams["istio-proxy-pod-ns"] = namespace
		istioIngressProxyCpu, err := getMetricPrometheusOCP("istio_proxies_cpu_custom", prometheusAPIMapParams)
		if err != nil {
			util.Log.Error(err)
			t.Error(err)
			t.FailNow()
		}
		util.Log.Info(istioIngressProxyCpu)
		util.Log.Info(" If istioIngressProxiesAcceptanceCpu is lower than ", istioIngressProxiesAcceptanceCpu)
	}
}

func TestIstioEgressProxiesMem(t *testing.T) {
	util.Log.Info("** TEST: TestIstioEgressProxiesMem")

	meshPods, err := getMeshEgressProxies()
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	for name, namespace := range meshPods {
		prometheusAPIMapParams["istio-proxy-pod-name"] = name
		prometheusAPIMapParams["istio-proxy-pod-ns"] = namespace
		istioEgressProxyMem, err := getMetricPrometheusOCP("istio_proxies_mem_custom", prometheusAPIMapParams)
		if err != nil {
			util.Log.Error(err)
			t.Error(err)
			t.FailNow()
		}
		util.Log.Info(istioEgressProxyMem)
		util.Log.Info(" If istioEgressProxiesAcceptanceMem is lower than ", istioEgressProxiesAcceptanceMem)
	}
}

func TestIstioEgressProxiesCpu(t *testing.T) {
	util.Log.Info("** TEST: TestIstioEgressProxiesCpu")

	meshPods, err := getMeshEgressProxies()
	if err != nil {
		util.Log.Error(err)
		t.Error(err)
		t.FailNow()
	}
	for name, namespace := range meshPods {
		prometheusAPIMapParams["istio-proxy-pod-name"] = name
		prometheusAPIMapParams["istio-proxy-pod-ns"] = namespace
		istioEgressProxyCpu, err := getMetricPrometheusOCP("istio_proxies_cpu_custom", prometheusAPIMapParams)
		if err != nil {
			util.Log.Error(err)
			t.Error(err)
			t.FailNow()
		}
		util.Log.Info(istioEgressProxyCpu)
		util.Log.Info(" If istioEgressProxiesAcceptanceCpu is lower than ", istioEgressProxiesAcceptanceCpu)
	}
}
