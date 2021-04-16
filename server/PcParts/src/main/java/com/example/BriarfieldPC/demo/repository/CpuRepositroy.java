package com.example.BriarfieldPC.demo.repository;

import com.example.BriarfieldPC.demo.entity.Cpu;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CpuRepositroy extends JpaRepository<Cpu,Long> {
    List<Cpu> findAll();
}
