package model_test

import (
	. "github.com/apache/servicecomb-rokie/pkg/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Kv mongodb service", func() {
	var s KVService
	var err error
	Describe("connecting db", func() {
		s, err = NewMongoService(Options{
			URI: "mongodb://rokie:123@127.0.0.1:27017",
		})
		It("should not return err", func() {
			Expect(err).Should(BeNil())
		})
	})

	Describe("put kv timeout", func() {
		Context("with labels app and service", func() {
			kv, err := s.CreateOrUpdate(&KV{
				Key:    "timeout",
				Value:  "2s",
				Domain: "default",
				Labels: map[string]string{
					"app":     "mall",
					"service": "cart",
				},
			})
			It("should not return err", func() {
				Expect(err).Should(BeNil())
			})
			It("should has revision", func() {
				Expect(kv.Revision).ShouldNot(BeZero())
			})
			It("should has ID", func() {
				Expect(kv.ID.Hex()).ShouldNot(BeEmpty())
			})

		})
		Context("with labels app, service and version", func() {
			kv, err := s.CreateOrUpdate(&KV{
				Key:    "timeout",
				Value:  "2s",
				Domain: "default",
				Labels: map[string]string{
					"app":     "mall",
					"service": "cart",
					"version": "1.0.0",
				},
			})
			oid, err := s.Exist("timeout", "default", map[string]string{
				"app":     "mall",
				"service": "cart",
				"version": "1.0.0",
			})
			It("should not return err", func() {
				Expect(err).Should(BeNil())
			})
			It("should has revision", func() {
				Expect(kv.Revision).ShouldNot(BeZero())
			})
			It("should has ID", func() {
				Expect(kv.ID.Hex()).ShouldNot(BeEmpty())
			})
			It("should exist", func() {
				Expect(oid).ShouldNot(BeEmpty())
			})
		})
		Context("with labels app,and update value", func() {
			beforeKV, err := s.CreateOrUpdate(&KV{
				Key:    "timeout",
				Value:  "1s",
				Domain: "default",
				Labels: map[string]string{
					"app": "mall",
				},
			})
			It("should not return err", func() {
				Expect(err).Should(BeNil())
			})
			kvs1, err := s.Find("default", WithKey("timeout"), WithLabels(map[string]string{
				"app": "mall",
			}), WithExactOne())
			log.Println("============", kvs1[0].Value)
			It("should be 1s", func() {
				Expect(kvs1[0].Value).Should(Equal(beforeKV.Value))
			})
			afterKV, err := s.CreateOrUpdate(&KV{
				Key:    "timeout",
				Value:  "3s",
				Domain: "default",
				Labels: map[string]string{
					"app": "mall",
				},
			})
			It("should has same id", func() {
				Expect(afterKV.ID.Hex()).Should(Equal(beforeKV.ID.Hex()))
			})
			oid, err := s.Exist("timeout", "default", map[string]string{
				"app": "mall",
			})
			It("should exists", func() {
				Expect(oid).Should(Equal(beforeKV.ID.Hex()))
			})
			kvs, err := s.Find("default", WithKey("timeout"), WithLabels(map[string]string{
				"app": "mall",
			}), WithExactOne())
			It("should be 3s", func() {
				Expect(kvs[0].Value).Should(Equal(afterKV.Value))
			})
		})
	})

	Describe("find kv timeout", func() {
		Context("with labels app", func() {
			kvs, err := s.Find("default", WithKey("timeout"), WithLabels(map[string]string{
				"app": "mall",
			}))
			It("should not return err", func() {
				Expect(err).Should(BeNil())
			})
			It("should has 3 records", func() {
				Expect(len(kvs)).Should(Equal(3))
			})

		})
	})
})
