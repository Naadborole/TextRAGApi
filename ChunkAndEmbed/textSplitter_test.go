package chunkandembed

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var txt string = `To help users get started, NVIDIA developed an AI workflow for retrieval-augmented generation. It includes a sample chatbot and the elements users need to create their own applications with this new method.

The workflow uses NVIDIA NeMo, a framework for developing and customizing generative AI models, as well as software like NVIDIA Triton Inference Server and NVIDIA TensorRT-LLM for running generative AI models in production.

The software components are all part of NVIDIA AI Enterprise, a software platform that accelerates development and deployment of production-ready AI with the security, support and stability businesses need.

Getting the best performance for RAG workflows requires massive amounts of memory and compute to move and process data. The NVIDIA GH200 Grace Hopper Superchip, with its 288GB of fast HBM3e memory and 8 petaflops of compute, is ideal — it can deliver a 150x speedup over using a CPU.

Once companies get familiar with RAG, they can combine a variety of off-the-shelf or custom LLMs with internal or external knowledge bases to create a wide range of assistants that help their employees and customers.

RAG doesn't require a data center. LLMs are debuting on Windows PCs, thanks to NVIDIA software that enables all sorts of applications users can access even on their laptops.`
	texts := make([]string, 1)
	texts[0] = txt
	var metadatas []map[string]any = make([]map[string]any, 0)
	metadatas = append(metadatas, make(map[string]any))
	metadatas[0]["source"] = "nvidia-blog"
	fmt.Println(SplitText(texts, metadatas))

}
